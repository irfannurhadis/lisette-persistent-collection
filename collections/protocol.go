package collections

import lisette "github.com/ivov/lisette/prelude"

type Collection[T any] interface {
	Count() int
	Empty() bool
	Reduce(T, func(T, T) T) T
}

type KvCollection[K comparable, V any] interface {
	ReduceKv(V, func(V, K, V) V) V
}

func (l List[T]) Count() int {
	return ListCount(l)
}

func (l List[T]) Empty() bool {
	if l.Tag == ListEmpty {
		return true
	}
	return false
}

func (l List[T]) Reduce(zero T, f func(T, T) T) T {
	acc := zero
	cur := l
	for {
		if cur.Tag == ListEmpty {
			break
		}
		tail := cur.Tail
		acc = f(acc, cur.Head)
		cur = *tail
	}
	return acc
}

func (v Vector[T]) Count() int {
	return v.N
}

func (v Vector[T]) Empty() bool {
	return v.N == 0
}

func (v Vector[T]) Reduce(zero T, f func(T, T) T) T {
	acc := zero
	i := 0
	for i < v.N {
		if x, ok := Vnth(v, i); ok {
			acc = f(acc, x)
		}
		i++
	}
	return acc
}

func (h HashMap[K, V]) Count() int {
	return h.N
}

func (h HashMap[K, V]) Empty() bool {
	return h.N == 0
}

func (h HashMap[K, V]) Reduce(zero V, f func(V, V) V) V {
	acc := zero
	for _, e := range Entries(h) {
		acc = f(acc, e.Value)
	}
	return acc
}

func (h HashSet[K]) Count() int {
	return h.Backing.N
}

func (h HashSet[K]) Empty() bool {
	return h.Backing.N == 0
}

func (h HashSet[K]) Reduce(zero K, f func(K, K) K) K {
	acc := zero
	for _, e := range Entries(h.Backing) {
		acc = f(acc, e.Key)
	}
	return acc
}

func (q Queue[T]) Reduce(zero T, f func(T, T) T) T {
	acc := zero
	cur := q.Front
	for {
		if cur.Tag == ListEmpty {
			break
		}
		tail := cur.Tail
		acc = f(acc, cur.Head)
		cur = *tail
	}
	rearRev := EmptyList[T]()
	cur = q.Rear
	for {
		if cur.Tag == ListEmpty {
			break
		}
		tail := cur.Tail
		rearRev = Cons(cur.Head, rearRev)
		cur = *tail
	}
	cur = rearRev
	for {
		if cur.Tag == ListEmpty {
			break
		}
		tail := cur.Tail
		acc = f(acc, cur.Head)
		cur = *tail
	}
	return acc
}

func TransduceVec[A any, R any](v Vector[A], stages []Stage[A], zero R, f func(R, A) R) R {
	acc := zero
	st := FreshState()
	i := 0
	for i < v.N {
		if st.Stopped {
			break
		}
		value := lisette.OptionFromCommaOk[A](Vnth(v, i))
		for _, stage := range stages {
			if value.IsNone() {
				break
			}
			tmp := stage.Apply(st, valueOf(value))
			next := tmp.First
			emitted := tmp.Second
			flag := tmp.Third
			st = next
			value = emitted
			if flag == MakeReducedStop() {
				copy2 := st
				copy2.Stopped = true
				st = copy2
				break
			}
		}
		if !value.IsNone() {
			acc = f(acc, valueOf(value))
		}
		i++
	}
	return acc
}

func TransduceList[A any, R any](xs List[A], stages []Stage[A], zero R, f func(R, A) R) R {
	acc := zero
	st := FreshState()
	cur := xs
	for {
		if cur.Tag == ListEmpty {
			break
		}
		tail := cur.Tail
		if st.Stopped {
			break
		}
		value := lisette.MakeOptionSome(cur.Head)
		for _, stage := range stages {
			if value.IsNone() {
				break
			}
			tmp := stage.Apply(st, valueOf(value))
			next := tmp.First
			emitted := tmp.Second
			flag := tmp.Third
			st = next
			value = emitted
			if flag == MakeReducedStop() {
				copy2 := st
				copy2.Stopped = true
				st = copy2
				break
			}
		}
		if !value.IsNone() {
			acc = f(acc, valueOf(value))
		}
		cur = *tail
	}
	return acc
}

func TransduceMapKeys[K comparable, R any](m HashMap[K, R], stages []Stage[K], zero R, f func(R, K) R) R {
	acc := zero
	st := FreshState()
	for _, e := range Entries(m) {
		if st.Stopped {
			break
		}
		value := lisette.MakeOptionSome(e.Key)
		for _, stage := range stages {
			if value.IsNone() {
				break
			}
			tmp := stage.Apply(st, valueOf(value))
			next := tmp.First
			emitted := tmp.Second
			flag := tmp.Third
			st = next
			value = emitted
			if flag == MakeReducedStop() {
				copy2 := st
				copy2.Stopped = true
				st = copy2
				break
			}
		}
		if !value.IsNone() {
			acc = f(acc, valueOf(value))
		}
	}
	return acc
}

func TransduceSet[K comparable, R any](s HashSet[K], stages []Stage[K], zero R, f func(R, K) R) R {
	acc := zero
	st := FreshState()
	for _, e := range Entries(s.Backing) {
		if st.Stopped {
			break
		}
		value := lisette.MakeOptionSome(e.Key)
		for _, stage := range stages {
			if value.IsNone() {
				break
			}
			tmp := stage.Apply(st, valueOf(value))
			next := tmp.First
			emitted := tmp.Second
			flag := tmp.Third
			st = next
			value = emitted
			if flag == MakeReducedStop() {
				copy2 := st
				copy2.Stopped = true
				st = copy2
				break
			}
		}
		if !value.IsNone() {
			acc = f(acc, valueOf(value))
		}
	}
	return acc
}

func Into[A any](stages []Stage[A], xs []A) Vector[A] {
	v := EmptyVector[A]()
	st := FreshState()
	for _, a := range xs {
		if st.Stopped {
			break
		}
		value := lisette.MakeOptionSome(a)
		for _, stage := range stages {
			if value.IsNone() {
				break
			}
			tmp := stage.Apply(st, valueOf(value))
			next := tmp.First
			emitted := tmp.Second
			flag := tmp.Third
			st = next
			value = emitted
			if flag == MakeReducedStop() {
				copy2 := st
				copy2.Stopped = true
				st = copy2
				break
			}
		}
		if !value.IsNone() {
			v = Vconj(v, valueOf(value))
		}
	}
	return v
}
