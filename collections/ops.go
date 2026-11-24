package collections

import "slices"

func Sunion[K comparable](a, b HashSet[K]) HashSet[K] {
	out := a
	for _, e := range Entries(b.Backing) {
		out = Sconj(out, e.Key)
	}
	return out
}

func Sintersection[K comparable](a, b HashSet[K]) HashSet[K] {
	out := EmptySet[K]()
	for _, e := range Entries(a.Backing) {
		if Scontains(b, e.Key) {
			out = Sconj(out, e.Key)
		}
	}
	return out
}

func Sdifference[K comparable](a, b HashSet[K]) HashSet[K] {
	out := EmptySet[K]()
	for _, e := range Entries(a.Backing) {
		if !Scontains(b, e.Key) {
			out = Sconj(out, e.Key)
		}
	}
	return out
}

func Mmerge[K comparable, V any](a, b HashMap[K, V]) HashMap[K, V] {
	out := a
	for _, e := range Entries(b) {
		out = Massoc(out, e.Key, e.Value)
	}
	return out
}

func MapKeys[K comparable, V any](m HashMap[K, V], f func(K) K) HashMap[K, V] {
	out := EmptyMap[K, V]()
	for _, e := range Entries(m) {
		out = Massoc(out, f(e.Key), e.Value)
	}
	return out
}

func MapValues[K comparable, V any](m HashMap[K, V], f func(V) V) HashMap[K, V] {
	out := EmptyMap[K, V]()
	for _, e := range Entries(m) {
		out = Massoc(out, e.Key, f(e.Value))
	}
	return out
}

func FilterKeys[K comparable, V any](m HashMap[K, V], p func(K) bool) HashMap[K, V] {
	out := EmptyMap[K, V]()
	for _, e := range Entries(m) {
		if p(e.Key) {
			out = Massoc(out, e.Key, e.Value)
		}
	}
	return out
}

func FilterValues[K comparable, V any](m HashMap[K, V], p func(V) bool) HashMap[K, V] {
	out := EmptyMap[K, V]()
	for _, e := range Entries(m) {
		if p(e.Value) {
			out = Massoc(out, e.Key, e.Value)
		}
	}
	return out
}

func EqSlice[A comparable](a, b []A) bool {
	return slices.Equal(a, b)
}
