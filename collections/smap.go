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

func visitKv[V any, R any](node *SNode[V], acc R, f func(R, string, V) R) R {
	out := acc
	for _, item := range node.Bucket {
		k := item.First
		v := item.Second
		out = f(out, k, v)
	}
	for _, child := range node.Kids {
		if child.Tag == lisette.OptionSome {
			out = visitKv(child.SomeVal, out, f)
		}
	}
	return out
}

func ReduceKv[V any, R any](m SMap[V], zero R, f func(R, string, V) R) R {
	return lisette.OptionMapOr(m.Root, zero, func(r *SNode[V]) R {
		return visitKv(r, zero, f)
	})
}

func FromEntries[V any](entries []lisette.Tuple2[string, V]) SMap[V] {
	m := EmptySmap[V]()
	for _, e := range entries {
		m = Sassoc(m, e.First, e.Second)
	}
	return m
}

func TransduceKeys[V any](m SMap[V], stages []Stage[string], zero int, f func(int, string) int) int {
	keys := ReduceKv(m, []string{}, func(acc []string, k string, _ V) []string {
		return append(slices.Clip(acc), k)
	})
	acc := zero
	st := FreshState()
	for _, k := range keys {
		if st.Stopped {
			break
		}
		value := lisette.MakeOptionSome(k)
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
