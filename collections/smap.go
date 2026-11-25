package collections

import (
	"slices"

	lisette "github.com/ivov/lisette/prelude"
)

const Smask int = 31

type SNode[V any] struct {
	Kids   []lisette.Option[*SNode[V]]
	Bucket []lisette.Tuple2[string, V]
}

type SMap[V any] struct {
	N    int
	Root lisette.Option[*SNode[V]]
}

func fnv(s string) uint64 {
	var h uint64 = 14695981039346656037
	for i := 0; i < len(s); i++ {
		b := s[i]
		h = (h ^ uint64(b)) * 1099511628211
	}
	return h
}

func EmptySmap[V any]() SMap[V] {
	return SMap[V]{N: 0}
}

func newSnode[V any]() SNode[V] {
	kids := []lisette.Option[*SNode[V]]{}
	for range 32 {
		kids = append(kids, lisette.MakeOptionNone[*SNode[V]]())
	}
	return SNode[V]{
		Kids:   kids,
		Bucket: []lisette.Tuple2[string, V]{},
	}
}

func Sget[V any](m SMap[V], k string) (V, bool) {
	if m.Root.Tag == lisette.OptionNone {
		var zero V
		return zero, false
	}
	r := m.Root.SomeVal
	h := fnv(k)
	node := r
	shift := 0
	for {
		for _, item := range node.Bucket {
			key := item.First
			v := item.Second
			if key == k {
				return v, true
			}
		}
		idx := int(h >> shift & uint64(Smask))
		subject := node.Kids[idx]
		if subject.Tag == lisette.OptionNone {
			var zero2 V
			return zero2, false
		}
		node = subject.SomeVal
		shift += 5
	}
}

func Sassoc[V any](m SMap[V], k string, v V) SMap[V] {
	h := fnv(k)
	return lisette.OptionMapOrElse(m.Root, func() SMap[V] {
		node := newSnode[V]()
		node.Bucket = append(node.Bucket, lisette.MakeTuple2[string, V](k, v))
		return SMap[V]{
			N:    m.N + 1,
			Root: lisette.MakeOptionSome(&node),
		}
	}, func(r *SNode[V]) SMap[V] {
		out := newSnode[V]()
		out.Kids = slices.Clone(r.Kids)
		out.Bucket = r.Bucket
		replaced := insertAt(&out, k, v, h, 0)
		var newN int
		if replaced {
			newN = m.N
		} else {
			newN = m.N + 1
		}
		return SMap[V]{
			N:    newN,
			Root: lisette.MakeOptionSome(&out),
		}
	})
}

func insertAt[V any](node *SNode[V], k string, v V, h uint64, shift int) bool {
	found := false
	bucket := []lisette.Tuple2[string, V]{}
	for _, item := range node.Bucket {
		key := item.First
		old := item.Second
		if key == k {
			bucket = append(bucket, lisette.MakeTuple2[string, V](k, v))
			found = true
		} else {
			bucket = append(bucket, lisette.MakeTuple2[string, V](key, old))
		}
	}
	if found {
		node.Bucket = bucket
		return true
	}
	if shift >= 60 {
		node.Bucket = append(slices.Clip(bucket), lisette.MakeTuple2[string, V](k, v))
		return false
	}
	idx := int(h >> shift & uint64(Smask))
	return lisette.OptionMapOrElse(node.Kids[idx], func() bool {
		leaf := newSnode[V]()
		leaf.Bucket = append(leaf.Bucket, lisette.MakeTuple2[string, V](k, v))
		node.Kids[idx] = lisette.MakeOptionSome(&leaf)
		return false
	}, func(child *SNode[V]) bool {
		next := newSnode[V]()
		next.Kids = slices.Clone(child.Kids)
		next.Bucket = child.Bucket
		replaced := insertAt(&next, k, v, h, shift+5)
		node.Kids[idx] = lisette.MakeOptionSome(&next)
		return replaced
	})
}

func Scount[V any](m SMap[V]) int {
	return m.N
}

func ReduceKv[V any, R any](m SMap[V], zero R, f func(R, string, V) R) R {
	acc := zero
	stack := []*SNode[V]{}
	if m.Root.Tag == lisette.OptionNone {
		return acc
	}
	stack = append(stack, m.Root.SomeVal)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = slices.Delete(stack, len(stack)-1, len(stack))
		for _, item := range node.Bucket {
			k := item.First
			v := item.Second
			acc = f(acc, k, v)
		}
		for _, child := range node.Kids {
			if child.Tag == lisette.OptionSome {
				stack = append(stack, child.SomeVal)
			}
		}
	}
	return acc
}

func FromEntries[V any](entries []lisette.Tuple2[string, V]) SMap[V] {
	m := EmptySmap[V]()
	for _, e := range entries {
		m = Sassoc(m, e.First, e.Second)
	}
	return m
}

type XAcc struct {
	Acc    int
	St     RState
	Stages []Stage[string]
	F      func(int, string) int
}

func visitKvFused[V any](node *SNode[V], xa *XAcc) {
	if xa.St.Stopped {
		return
	}
	for _, item := range node.Bucket {
		k := item.First
		if xa.St.Stopped {
			return
		}
		value := lisette.MakeOptionSome(k)
		for _, stage := range xa.Stages {
			if value.IsNone() {
				break
			}
			arg := xa.St
			tmp := stage.Apply(arg, valueOf(value))
			next := tmp.First
			emitted := tmp.Second
			flag := tmp.Third
			xa.St = next
			value = emitted
			if flag == MakeReducedStop() {
				copy2 := xa.St
				copy2.Stopped = true
				xa.St = copy2
				break
			}
		}
		if !value.IsNone() {
			callee := xa.F
			arg2 := xa.Acc
			xa.Acc = callee(arg2, valueOf(value))
		}
	}
	for _, child := range node.Kids {
		if xa.St.Stopped {
			return
		}
		if child.Tag == lisette.OptionSome {
			visitKvFused(child.SomeVal, xa)
		}
	}
}

func TransduceKeys[V any](m SMap[V], stages []Stage[string], zero int, f func(int, string) int) int {
	xa := XAcc{
		Acc:    zero,
		St:     FreshState(),
		Stages: stages,
		F:      f,
	}
	if m.Root.Tag == lisette.OptionSome {
		visitKvFused(m.Root.SomeVal, &xa)
	}
	return xa.Acc
}
