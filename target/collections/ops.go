package collections

import "slices"

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:9:8
func Sunion[K comparable](a, b HashSet[K]) HashSet[K] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:10:5
	out := a
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:11:5
	for _, e := range b.Backing.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:11:34
		out = Sconj(out, e.Key)
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:12:5
	return out
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:15:8
func Sintersection[K comparable](a, b HashSet[K]) HashSet[K] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:16:5
	out := EmptySet[K]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:17:5
	for _, e := range a.Backing.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:18:9
		if Scontains(b, e.Key) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:18:34
			out = Sconj(out, e.Key)
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:20:5
	return out
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:23:8
func Sdifference[K comparable](a, b HashSet[K]) HashSet[K] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:24:5
	out := EmptySet[K]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:25:5
	for _, e := range a.Backing.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:26:9
		if !Scontains(b, e.Key) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:26:35
			out = Sconj(out, e.Key)
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:28:5
	return out
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:36:8
func Mmerge[K comparable, V any](a, b HashMap[K, V]) HashMap[K, V] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:37:5
	out := a
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:38:5
	for _, e := range b.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:38:26
		out = Massoc(out, e.Key, e.Value)
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:39:5
	return out
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:44:8
func MapKeys[K comparable, V any](m HashMap[K, V], f func(K) K) HashMap[K, V] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:45:5
	out := EmptyMap[K, V]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:46:5
	for _, e := range m.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:46:26
		out = Massoc(out, f(e.Key), e.Value)
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:47:5
	return out
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:50:8
func MapValues[K comparable, V any](m HashMap[K, V], f func(V) V) HashMap[K, V] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:51:5
	out := EmptyMap[K, V]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:52:5
	for _, e := range m.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:52:26
		out = Massoc(out, e.Key, f(e.Value))
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:53:5
	return out
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:56:8
func FilterKeys[K comparable, V any](m HashMap[K, V], p func(K) bool) HashMap[K, V] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:57:5
	out := EmptyMap[K, V]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:58:5
	for _, e := range m.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:58:26
		if p(e.Key) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:58:40
			out = Massoc(out, e.Key, e.Value)
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:59:5
	return out
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:62:8
func FilterValues[K comparable, V any](m HashMap[K, V], p func(V) bool) HashMap[K, V] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:63:5
	out := EmptyMap[K, V]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:64:5
	for _, e := range m.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:64:26
		if p(e.Value) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:64:42
			out = Massoc(out, e.Key, e.Value)
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:65:5
	return out
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:68:8
func EqSlice[A comparable](a, b []A) bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.lis:68:68
	return slices.Equal(a, b)
}
