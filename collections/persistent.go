package collections

import (
	"fmt"
	"slices"

	lisette "github.com/ivov/lisette/prelude"
)

func MakeListEmpty[T any]() List[T] {
	return List[T]{Tag: ListEmpty}
}

func MakeListCons[T any](arg0 T, arg1 *List[T]) List[T] {
	return List[T]{Tag: ListCons, Head: arg0, Tail: arg1}
}

func MakeUnitValOnly() UnitVal {
	return UnitVal{Tag: UnitValOnly}
}

type ListTag uint8

const (
	ListEmpty ListTag = iota
	ListCons
)

type List[T any] struct {
	Tag  ListTag
	Head T
	Tail *List[T]
}

func EmptyList[T any]() List[T] {
	return MakeListEmpty[T]()
}

func Cons[T any](head T, tail List[T]) List[T] {
	return List[T]{
		Tag:  ListCons,
		Head: head,
		Tail: &tail,
	}
}

func First[T any](xs List[T]) (T, bool) {
	if xs.Tag == ListEmpty {
		var zero T
		return zero, false
	}
	return xs.Head, true
}

func Rest[T any](xs List[T]) List[T] {
	if xs.Tag == ListEmpty {
		return MakeListEmpty[T]()
	}
	return *xs.Tail
}

func ListCount[T any](xs List[T]) int {
	n := 0
	cur := xs
	for {
		if cur.Tag == ListEmpty {
			return n
		}
		tail := cur.Tail
		n++
		cur = *tail
	}
}

func FromSlice[T any](xs []T) List[T] {
	out := EmptyList[T]()
	i := len(xs) - 1
	for i >= 0 {
		out = Cons(xs[i], out)
		i--
	}
	return out
}

const Bits int = 5

const Width int = 32

const Mask int = 31

type Vector[T any] struct {
	N      int
	Chunks [][]T
	Tail   []T
}

func EmptyVector[T any]() Vector[T] {
	return Vector[T]{
		N:      0,
		Chunks: [][]T{},
		Tail:   []T{},
	}
}

func Vconj[T any](v Vector[T], x T) Vector[T] {
	tail := append(slices.Clip(v.Tail), x)
	if len(tail) < Width {
		return Vector[T]{
			N:      v.N + 1,
			Chunks: v.Chunks,
			Tail:   tail,
		}
	}
	chunks := [][]T{}
	for _, c := range v.Chunks {
		chunks = append(chunks, c)
	}
	chunks = append(chunks, tail)
	return Vector[T]{
		N:      v.N + 1,
		Chunks: chunks,
		Tail:   []T{},
	}
}

func VtoSlice[T any](v Vector[T]) []T {
	out := []T{}
	for _, c := range v.Chunks {
		for _, x := range c {
			out = append(out, x)
		}
	}
	for _, x := range v.Tail {
		out = append(out, x)
	}
	return out
}

func Vnth[T any](v Vector[T], i int) (T, bool) {
	if i < 0 || i >= v.N {
		var zero2 T
		return zero2, false
	}
	chunk := i / Width
	off := i % Width
	if chunk < len(v.Chunks) {
		v2 := lisette.SliceGet(v.Chunks[chunk], off)
		if v2.Tag == lisette.OptionSome {
			return v2.SomeVal, true
		}
		var zero3 T
		return zero3, false
	}
	v3 := lisette.SliceGet(v.Tail, i-len(v.Chunks)*Width)
	if v3.Tag == lisette.OptionSome {
		return v3.SomeVal, true
	}
	var zero T
	return zero, false
}

type Entry[K any, V any] struct {
	Key   K
	Value V
}

type HNode[K any, V any] struct {
	Kids   []lisette.Option[*HNode[K, V]]
	Bucket []Entry[K, V]
}

type HashMap[K comparable, V any] struct {
	N    int
	Root lisette.Option[*HNode[K, V]]
}

func khash[K comparable](k K) uint64 {
	text := fmt.Sprint(k)
	var h uint64 = 14695981039346656037
	for i := 0; i < len(text); i++ {
		b := text[i]
		h = (h ^ uint64(b)) * 1099511628211
	}
	return h
}

func newHnode[K comparable, V any]() HNode[K, V] {
	kids := []lisette.Option[*HNode[K, V]]{}
	for range 32 {
		kids = append(kids, lisette.MakeOptionNone[*HNode[K, V]]())
	}
	return HNode[K, V]{
		Kids:   kids,
		Bucket: []Entry[K, V]{},
	}
}

func EmptyMap[K comparable, V any]() HashMap[K, V] {
	return HashMap[K, V]{N: 0}
}

func Mget[K comparable, V any](m HashMap[K, V], k K) (V, bool) {
	if m.Root.Tag == lisette.OptionNone {
		var zero V
		return zero, false
	}
	r := m.Root.SomeVal
	h := khash(k)
	node := r
	shift := 0
	for {
		for _, e := range node.Bucket {
			if e.Key == k {
				return e.Value, true
			}
		}
		idx := int(h >> shift & 31)
		subject := node.Kids[idx]
		if subject.Tag == lisette.OptionNone {
			var zero2 V
			return zero2, false
		}
		node = subject.SomeVal
		shift += 5
	}
}

func hput[K comparable, V any](node *HNode[K, V], k K, v V, h uint64, shift int) bool {
	found := false
	bucket := []Entry[K, V]{}
	for _, e := range node.Bucket {
		if e.Key == k {
			bucket = append(bucket, Entry[K, V]{
				Key:   k,
				Value: v,
			})
			found = true
		} else {
			bucket = append(bucket, e)
		}
	}
	if found {
		node.Bucket = bucket
		return false
	}
	idx := int(h >> shift & 31)
	return lisette.OptionMapOrElse(node.Kids[idx], func() bool {
		leaf := newHnode[K, V]()
		leaf.Bucket = append(leaf.Bucket, Entry[K, V]{
			Key:   k,
			Value: v,
		})
		node.Kids[idx] = lisette.MakeOptionSome(&leaf)
		return false
	}, func(child *HNode[K, V]) bool {
		next := newHnode[K, V]()
		next.Kids = slices.Clone(child.Kids)
		next.Bucket = child.Bucket
		replaced := hput(&next, k, v, h, shift+5)
		node.Kids[idx] = lisette.MakeOptionSome(&next)
		return replaced
	})
}

func Massoc[K comparable, V any](m HashMap[K, V], k K, v V) HashMap[K, V] {
	return lisette.OptionMapOrElse(m.Root, func() HashMap[K, V] {
		node := newHnode[K, V]()
		node.Bucket = append(node.Bucket, Entry[K, V]{
			Key:   k,
			Value: v,
		})
		return HashMap[K, V]{
			N:    m.N + 1,
			Root: lisette.MakeOptionSome(&node),
		}
	}, func(r *HNode[K, V]) HashMap[K, V] {
		out := newHnode[K, V]()
		out.Kids = slices.Clone(r.Kids)
		out.Bucket = r.Bucket
		arg := &out
		replaced := hput(arg, k, v, khash(k), 0)
		var newN int
		if replaced {
			newN = m.N
		} else {
			newN = m.N + 1
		}
		return HashMap[K, V]{
			N:    newN,
			Root: lisette.MakeOptionSome(&out),
		}
	})
}

func hdel[K comparable, V any](node *HNode[K, V], k K, h uint64, shift int) bool {
	found := false
	bucket := []Entry[K, V]{}
	for _, e := range node.Bucket {
		if e.Key == k {
			found = true
		} else {
			bucket = append(bucket, e)
		}
	}
	if found {
		node.Bucket = bucket
		return true
	}
	idx := int(h >> shift & 31)
	return lisette.OptionMapOr(node.Kids[idx], false, func(child *HNode[K, V]) bool {
		next := newHnode[K, V]()
		next.Kids = slices.Clone(child.Kids)
		next.Bucket = child.Bucket
		removed := hdel(&next, k, h, shift+5)
		node.Kids[idx] = lisette.MakeOptionSome(&next)
		return removed
	})
}

func Mdissoc[K comparable, V any](m HashMap[K, V], k K) HashMap[K, V] {
	return lisette.OptionMapOr(m.Root, m, func(r *HNode[K, V]) HashMap[K, V] {
		out := newHnode[K, V]()
		out.Kids = slices.Clone(r.Kids)
		out.Bucket = r.Bucket
		arg := &out
		if hdel(arg, k, khash(k), 0) {
			return HashMap[K, V]{
				N:    m.N - 1,
				Root: lisette.MakeOptionSome(&out),
			}
		}
		return m
	})
}

func walkEntries[K comparable, V any](node *HNode[K, V], out []Entry[K, V]) []Entry[K, V] {
	acc := out
	for _, e := range node.Bucket {
		acc = append(acc, e)
	}
	for _, child := range node.Kids {
		if child.Tag == lisette.OptionSome {
			acc = walkEntries(child.SomeVal, acc)
		}
	}
	return acc
}

func Entries[K comparable, V any](m HashMap[K, V]) []Entry[K, V] {
	return lisette.OptionMapOrElse(m.Root, func() []Entry[K, V] {
		return []Entry[K, V]{}
	}, func(r *HNode[K, V]) []Entry[K, V] {
		return walkEntries(r, []Entry[K, V]{})
	})
}

type UnitValTag uint8

const (
	UnitValOnly UnitValTag = iota
)

type UnitVal struct {
	Tag UnitValTag
}

type HashSet[T comparable] struct {
	Backing HashMap[T, UnitVal]
}

func EmptySet[T comparable]() HashSet[T] {
	return HashSet[T]{Backing: EmptyMap[T, UnitVal]()}
}

func Sconj[T comparable](s HashSet[T], x T) HashSet[T] {
	return HashSet[T]{Backing: Massoc(s.Backing, x, MakeUnitValOnly())}
}

func Scontains[T comparable](s HashSet[T], x T) bool {
	_, ok := Mget(s.Backing, x)
	return ok
}

type Queue[T any] struct {
	Front List[T]
	Rear  List[T]
}

func EmptyQueue[T any]() Queue[T] {
	return Queue[T]{
		Front: EmptyList[T](),
		Rear:  EmptyList[T](),
	}
}

func reverseList[T any](xs, acc List[T]) List[T] {
	if xs.Tag == ListEmpty {
		return acc
	}
	arg := *xs.Tail
	return reverseList(arg, Cons(xs.Head, acc))
}

func normalize[T any](q Queue[T]) Queue[T] {
	if _, ok := First(q.Front); ok {
		return q
	}
	return Queue[T]{
		Front: reverseList(q.Rear, EmptyList[T]()),
		Rear:  EmptyList[T](),
	}
}

func Qconj[T any](q Queue[T], x T) Queue[T] {
	return normalize(Queue[T]{
		Front: q.Front,
		Rear:  Cons(x, q.Rear),
	})
}

func Qpeek[T any](q Queue[T]) (T, bool) {
	return First(q.Front)
}

func Qpop[T any](q Queue[T]) Queue[T] {
	return normalize(Queue[T]{
		Front: Rest(q.Front),
		Rear:  q.Rear,
	})
}
