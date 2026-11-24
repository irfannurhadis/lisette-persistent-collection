package collections

import (
	"fmt"
	"slices"

	lisette "github.com/ivov/lisette/prelude"
)

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:19:5
func MakeReducedContinue() Reduced {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:19:5
	return Reduced{Tag: ReducedContinue}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:20:5
func MakeReducedStop() Reduced {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:20:5
	return Reduced{Tag: ReducedStop}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:17:10
type ReducedTag uint8

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:17:10
const (
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:17:10
	ReducedContinue ReducedTag = iota
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:17:10
	ReducedStop
)

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:17:10
type Reduced struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:17:10
	Tag ReducedTag
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:17:10
func (r Reduced) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:17:10
	switch r.Tag {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:17:10
	case ReducedContinue:
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:17:10
		return "Continue"
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:17:10
	case ReducedStop:
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:17:10
		return "Stop"
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:17:10
	default:
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:17:10
		return fmt.Sprintf("Reduced(%d)", r.Tag)
	}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:23:12
type RState struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:23:12
	Cursor int
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:23:12
	Taken int
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:23:12
	Stopped bool
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:23:12
func (r RState) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:23:12
	return fmt.Sprintf("RState { cursor: %s, taken: %s, stopped: %s }", lisette.Debug(r.Cursor), lisette.Debug(r.Taken), lisette.Debug(r.Stopped))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:29:8
func FreshState() RState {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:30:5
	return RState{}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:33:4
func valueOf[A any](v lisette.Option[A]) A {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:34:5
	if v.Tag != lisette.OptionSome {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:34:28
		panic("transducer: absent value")
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:34:5
	x := v.SomeVal
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:35:5
	return x
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:40:12
type Stage[A any] struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:40:12
	Apply func(RState, A) lisette.Tuple3[RState, lisette.Option[A], Reduced]
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:40:12
func (s Stage[A]) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:40:12
	return fmt.Sprintf("Stage { apply: %p }", s.Apply)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:46:12
type Pipeline[A any] struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:46:12
	Stages []Stage[A]
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:46:12
func (p Pipeline[A]) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:46:12
	return fmt.Sprintf("Pipeline { stages: %s }", lisette.Debug(p.Stages))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:50:8
func Identity[A any]() Pipeline[A] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:51:5
	return Pipeline[A]{Stages: /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:52:17*/ []Stage[A]{{Apply: /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:53:20*/ func(st RState, a A) lisette.Tuple3[RState, lisette.Option[A], Reduced] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:53:39
		return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](st, lisette.MakeOptionSome(a), MakeReducedContinue())
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:53:70*/}}}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:58:8
func Map[A any](f func(A) A) Pipeline[A] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:59:5
	return Pipeline[A]{Stages: /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:60:17*/ []Stage[A]{{Apply: /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:61:20*/ func(st RState, a A) lisette.Tuple3[RState, lisette.Option[A], Reduced] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:61:39
		return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](st, lisette.MakeOptionSome(f(a)), MakeReducedContinue())
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:61:73*/}}}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:66:8
func Filter[A any](p func(A) bool) Pipeline[A] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:67:5
	return Pipeline[A]{Stages: /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:68:17*/ []Stage[A]{{Apply: /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:69:20*/ func(st RState, a A) lisette.Tuple3[RState, lisette.Option[A], Reduced] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:70:17
		if p(a) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:70:27
			return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](st, lisette.MakeOptionSome(a), MakeReducedContinue())
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:71:24
		return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](st, lisette.MakeOptionNone[A](), MakeReducedContinue())
	}}}}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:77:8
func Take[A any](n int) Pipeline[A] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:78:5
	return Pipeline[A]{Stages: /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:79:17*/ []Stage[A]{{Apply: /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:80:20*/ func(st RState, a A) lisette.Tuple3[RState, lisette.Option[A], Reduced] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:81:17
		if st.Taken >= n {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:81:36
			copy2 := st
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:81:36
			copy2.Stopped = true
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:81:36
			return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](copy2, lisette.MakeOptionNone[A](), MakeReducedStop())
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:82:24
		st.Taken = st.Taken + 1
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:82:24
		return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](st, lisette.MakeOptionSome(a), MakeReducedContinue())
	}}}}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:88:8
func Drop[A any](n int) Pipeline[A] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:89:5
	return Pipeline[A]{Stages: /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:90:17*/ []Stage[A]{{Apply: /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:91:20*/ func(st RState, a A) lisette.Tuple3[RState, lisette.Option[A], Reduced] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:92:17
		if st.Cursor < n {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:92:36
			copy2 := st
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:92:36
			copy2.Cursor = st.Cursor + 1
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:92:36
			return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](copy2, lisette.MakeOptionNone[A](), MakeReducedContinue())
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:93:24
		return lisette.MakeTuple3[RState, lisette.Option[A], Reduced](st, lisette.MakeOptionSome(a), MakeReducedContinue())
	}}}}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:99:8
func Compose[A any](first, second Pipeline[A]) Pipeline[A] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:100:5
	stages := []Stage[A]{}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:101:5
	for _, s := range first.Stages {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:101:29
		stages = append(stages, s)
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:102:5
	for _, s := range second.Stages {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:102:30
		stages = append(stages, s)
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:103:5
	return Pipeline[A]{Stages: stages}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:110:4
func reduce[A any, R any](stages []Stage[A], xs []A, zero R, step func(R, A) R) R {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:111:5
	acc := zero
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:112:5
	state := FreshState()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:113:5
	for _, a := range xs {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:114:9
		if state.Stopped {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:114:28
			break
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:115:9
		value := lisette.MakeOptionSome(a)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:116:9
		for _, stage := range stages {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:117:13
			if value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:117:34
				break
			}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:118:13
			tmp := stage.Apply(state, valueOf(value))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:118:13
			nextState := tmp.First
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:118:13
			emitted := tmp.Second
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:118:13
			flag := tmp.Third
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:119:13
			state = nextState
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:120:13
			value = emitted
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:121:13
			if flag == MakeReducedStop() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:122:17
				copy2 := state
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:122:17
				copy2.Stopped = true
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:122:17
				state = copy2
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:123:17
				break
			}
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:126:9
		if !value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:126:31
			acc = step(acc, valueOf(value))
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:128:5
	return acc
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:131:8
func Run[A any](pipeline Pipeline[A], xs []A) []A {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:132:5
	return reduce(pipeline.Stages, xs, []A{}, func(acc []A, a A) []A {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:132:71
		return append(slices.Clip(acc), a)
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:132:84*/)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:135:8
func Fold[A any, R any](pipeline Pipeline[A], xs []A, zero R, f func(R, A) R) R {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:136:5
	return reduce(pipeline.Stages, xs, zero, f)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:139:8
func Find[A any](pipeline Pipeline[A], xs []A, p func(A) bool) (A, bool) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:140:5
	state := FreshState()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:141:5
	for _, a := range xs {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:142:9
		if state.Stopped {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:142:28
			break
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:143:9
		value := lisette.MakeOptionSome(a)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:144:9
		for _, stage := range pipeline.Stages {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:145:13
			if value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:145:34
				break
			}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:146:13
			tmp := stage.Apply(state, valueOf(value))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:146:13
			nextState := tmp.First
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:146:13
			emitted := tmp.Second
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:146:13
			flag := tmp.Third
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:147:13
			state = nextState
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:148:13
			value = emitted
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:149:13
			if flag == MakeReducedStop() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:150:17
				copy2 := state
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:150:17
				copy2.Stopped = true
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:150:17
				state = copy2
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:151:17
				break
			}
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:154:9
		if !value.IsNone() && p(valueOf(value)) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:154:53
			if value.Tag == lisette.OptionSome {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:154:53
				return value.SomeVal, true
			}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:154:53
			var zero2 A
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:154:53
			return zero2, false
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:156:5
	var zero A
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:156:5
	return zero, false
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:165:8
func MapCollect[A any, B any](xs []A, f func(A) B) []B {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:166:5
	out := []B{}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:167:5
	for _, a := range xs {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:167:19
		out = append(out, f(a))
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:168:5
	return out
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:171:8
func FilterCollect[A any](xs []A, p func(A) bool) []A {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:172:5
	return Run(Filter(p), xs)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:175:8
func TakeCollect[A any](xs []A, n int) []A {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:176:5
	return Run(Take[A](n), xs)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:179:8
func DropCollect[A any](xs []A, n int) []A {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:180:5
	return Run(Drop[A](n), xs)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:183:8
func SumCollect(xs []int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:184:5
	return Fold(Identity[int](), xs, 0, func(acc, a int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:184:48
		return acc + a
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:184:55*/)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:188:8
func SelectIds(ids []string, keep func(string) bool, limit int) []string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:189:5
	return Run(Compose(Filter(keep), Take[string](limit)), ids)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:192:8
func CountMatching[A any](xs []A, p func(A) bool) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:193:5
	return Fold(Filter(p), xs, 0, func(acc int, _ A) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:193:46
		return acc + 1
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.lis:193:53*/)
}
