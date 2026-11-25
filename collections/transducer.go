package collections

import (
	"slices"

	lisette "github.com/ivov/lisette/prelude"
)

func MakeReducedContinue() Reduced {
	return Reduced{Tag: ReducedContinue}
}

func MakeReducedStop() Reduced {
	return Reduced{Tag: ReducedStop}
}

type ReducedTag uint8

const (
	ReducedContinue ReducedTag = iota
	ReducedStop
)

type Reduced struct {
	Tag ReducedTag
}

type RState struct {
	Cursor  int
	Taken   int
	Stopped bool
}

func FreshState() RState {
	return RState{}
}

func valueOf[A any](v lisette.Option[A]) A {
	if v.Tag != lisette.OptionSome {
		panic("transducer: absent value")
	}
	x := v.SomeVal
	return x
}

type Stage[A any] struct {
	Apply func(RState, A) lisette.Tuple3[RState, lisette.Option[A], Reduced]
}

type Pipeline[A any] struct {
	Stages []Stage[A]
}

func Identity[A any]() Pipeline[A] {
	return Pipeline[A]{Stages: []Stage[A]{{Apply: func(st RState, a A) lisette.Tuple3[RState, lisette.Option[A], Reduced] {
		return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](st, lisette.MakeOptionSome(a), MakeReducedContinue())
	}}}}
}

func Map[A any](f func(A) A) Pipeline[A] {
	return Pipeline[A]{Stages: []Stage[A]{{Apply: func(st RState, a A) lisette.Tuple3[RState, lisette.Option[A], Reduced] {
		return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](st, lisette.MakeOptionSome(f(a)), MakeReducedContinue())
	}}}}
}

func Filter[A any](p func(A) bool) Pipeline[A] {
	return Pipeline[A]{Stages: []Stage[A]{{Apply: func(st RState, a A) lisette.Tuple3[RState, lisette.Option[A], Reduced] {
		if p(a) {
			return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](st, lisette.MakeOptionSome(a), MakeReducedContinue())
		}
		return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](st, lisette.MakeOptionNone[A](), MakeReducedContinue())
	}}}}
}

func Take[A any](n int) Pipeline[A] {
	return Pipeline[A]{Stages: []Stage[A]{{Apply: func(st RState, a A) lisette.Tuple3[RState, lisette.Option[A], Reduced] {
		if n <= 0 {
			copy2 := st
			copy2.Stopped = true
			return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](copy2, lisette.MakeOptionNone[A](), MakeReducedStop())
		}
		nextTaken := st.Taken + 1
		if nextTaken >= n {
			copy3 := st
			copy3.Taken = nextTaken
			copy3.Stopped = true
			return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](copy3, lisette.MakeOptionSome(a), MakeReducedStop())
		}
		st.Taken = nextTaken
		return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](st, lisette.MakeOptionSome(a), MakeReducedContinue())
	}}}}
}

func Drop[A any](n int) Pipeline[A] {
	return Pipeline[A]{Stages: []Stage[A]{{Apply: func(st RState, a A) lisette.Tuple3[RState, lisette.Option[A], Reduced] {
		if st.Cursor < n {
			copy2 := st
			copy2.Cursor = st.Cursor + 1
			return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](copy2, lisette.MakeOptionNone[A](), MakeReducedContinue())
		}
		return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](st, lisette.MakeOptionSome(a), MakeReducedContinue())
	}}}}
}

func Compose[A any](first, second Pipeline[A]) Pipeline[A] {
	stages := []Stage[A]{}
	for _, s := range first.Stages {
		stages = append(stages, s)
	}
	for _, s := range second.Stages {
		stages = append(stages, s)
	}
	return Pipeline[A]{Stages: stages}
}

func reduce[A any, R any](stages []Stage[A], xs []A, zero R, step func(R, A) R) R {
	acc := zero
	state := FreshState()
	for _, a := range xs {
		if state.Stopped {
			break
		}
		value := lisette.MakeOptionSome(a)
		for _, stage := range stages {
			if value.IsNone() {
				break
			}
			tmp := stage.Apply(state, valueOf(value))
			nextState := tmp.First
			emitted := tmp.Second
			flag := tmp.Third
			state = nextState
			value = emitted
			if flag == MakeReducedStop() {
				copy2 := state
				copy2.Stopped = true
				state = copy2
				break
			}
		}
		if !value.IsNone() {
			acc = step(acc, valueOf(value))
		}
	}
	return acc
}

func Run[A any](pipeline Pipeline[A], xs []A) []A {
	return reduce(pipeline.Stages, xs, []A{}, func(acc []A, a A) []A {
		return append(slices.Clip(acc), a)
	})
}

func Fold[A any, R any](pipeline Pipeline[A], xs []A, zero R, f func(R, A) R) R {
	return reduce(pipeline.Stages, xs, zero, f)
}

func Find[A any](pipeline Pipeline[A], xs []A, p func(A) bool) (A, bool) {
	state := FreshState()
	for _, a := range xs {
		if state.Stopped {
			break
		}
		value := lisette.MakeOptionSome(a)
		for _, stage := range pipeline.Stages {
			if value.IsNone() {
				break
			}
			tmp := stage.Apply(state, valueOf(value))
			nextState := tmp.First
			emitted := tmp.Second
			flag := tmp.Third
			state = nextState
			value = emitted
			if flag == MakeReducedStop() {
				copy2 := state
				copy2.Stopped = true
				state = copy2
				break
			}
		}
		if !value.IsNone() && p(valueOf(value)) {
			if value.Tag == lisette.OptionSome {
				return value.SomeVal, true
			}
			var zero2 A
			return zero2, false
		}
	}
	var zero A
	return zero, false
}

func MapCollect[A any, B any](xs []A, f func(A) B) []B {
	out := []B{}
	for _, a := range xs {
		out = append(out, f(a))
	}
	return out
}

func FilterCollect[A any](xs []A, p func(A) bool) []A {
	return Run(Filter(p), xs)
}

func TakeCollect[A any](xs []A, n int) []A {
	return Run(Take[A](n), xs)
}

func DropCollect[A any](xs []A, n int) []A {
	return Run(Drop[A](n), xs)
}

func SumCollect(xs []int) int {
	return Fold(Identity[int](), xs, 0, func(acc, a int) int {
		return acc + a
	})
}

func SelectIds(ids []string, keep func(string) bool, limit int) []string {
	return Run(Compose(Filter(keep), Take[string](limit)), ids)
}

func CountMatching[A any](xs []A, p func(A) bool) int {
	return Fold(Filter(p), xs, 0, func(acc int, _ A) int {
		return acc + 1
	})
}
