package collections

import (
	"fmt"
	"slices"

	lisette "github.com/ivov/lisette/prelude"
)

const Smask int = 31

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:8:12
type SNode[V any] struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:8:12
	Kids []lisette.Option[*SNode[V]]
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:8:12
	Bucket []lisette.Tuple2[string, V]
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:8:12
func (s SNode[V]) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:8:12
	return fmt.Sprintf("SNode { kids: %s, bucket: %s }", lisette.Debug(s.Kids), lisette.Debug(s.Bucket))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:15:12
type SMap[V any] struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:15:12
	N int
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:15:12
	Root lisette.Option[*SNode[V]]
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:15:12
func (s SMap[V]) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:15:12
	return fmt.Sprintf("SMap { n: %s, root: %s }", lisette.Debug(s.N), lisette.Debug(s.Root))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:20:4
func fnv(s string) uint64 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:21:5
	var h uint64 = 14695981039346656037
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:22:5
	for i := 0; i < len(s); i++ {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:22:5
		b := s[i]
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:23:9
		h = (h ^ uint64(b)) * 1099511628211
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:25:5
	return h
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:28:8
func EmptySmap[V any]() SMap[V] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:28:37
	return SMap[V]{N: 0}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:30:4
func newSnode[V any]() SNode[V] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:31:5
	kids := []lisette.Option[*SNode[V]]{}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:32:5
	for range 32 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:32:22
		kids = append(kids, lisette.MakeOptionNone[*SNode[V]]())
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:33:5
	return SNode[V]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:33:13
		Kids: kids,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:33:19
		Bucket: []lisette.Tuple2[string, V]{},
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:33:53*/
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:36:8
func Sget[V any](m SMap[V], k string) (V, bool) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:37:5
	if m.Root.Tag == lisette.OptionNone {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:38:17
		var zero V
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:38:17
		return zero, false
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:39:20
	r := m.Root.SomeVal
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:40:13
	h := fnv(k)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:41:13
	node := r
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:42:13
	shift := 0
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:43:13
	for {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:44:17
		for _, item := range node.Bucket {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:44:17
			key := item.First
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:44:17
			v := item.Second
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:45:21
			if key == k {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:45:35
				return v, true
			}
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:47:17
		idx := int(h >> shift & uint64(Smask))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:48:17
		subject := node.Kids[idx]
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:48:17
		if subject.Tag == lisette.OptionNone {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:50:29
			var zero2 V
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:50:29
			return zero2, false
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:49:38
		node = subject.SomeVal
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:49:52
		shift += 5
	}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:57:8
func Sassoc[V any](m SMap[V], k string, v V) SMap[V] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:58:5
	h := fnv(k)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:59:5
	return lisette.OptionMapOrElse(m.Root /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:60:9*/, func() SMap[V] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:61:13
		node := newSnode[V]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:62:13
		node.Bucket = append(node.Bucket, lisette.MakeTuple2[string, V](k, v))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:63:13
		return SMap[V]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:63:20
			N: m.N + 1,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:63:32
			Root: lisette.MakeOptionSome(&node),
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:63:51*/
	}, /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:65:9*/ func(r *SNode[V]) SMap[V] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:66:13
		out := newSnode[V]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:67:13
		out.Kids = slices.Clone(r.Kids)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:68:13
		out.Bucket = r.Bucket
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:69:13
		insertAt(&out, k, v, h, 0)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:70:13
		return SMap[V]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:70:20
			N: m.N + 1,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:70:32
			Root: lisette.MakeOptionSome(&out),
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:70:50*/
	})
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:75:4
func insertAt[V any](node *SNode[V], k string, v V, h uint64, shift int) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:77:5
	found := false
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:78:5
	bucket := []lisette.Tuple2[string, V]{}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:79:5
	for _, item := range node.Bucket {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:79:5
		key := item.First
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:79:5
		old := item.Second
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:80:9
		if key == k {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:80:23
			bucket = append(bucket, lisette.MakeTuple2[string, V](k, v))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:80:55
			found = true
		} else {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:81:16
			bucket = append(bucket, lisette.MakeTuple2[string, V](key, old))
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:83:5
	if found {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:83:16
		node.Bucket = bucket
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:83:38
		return
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:84:5
	idx := int(h >> shift & uint64(Smask))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:85:5
	subject := node.Kids[idx]
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:85:5
	if subject.Tag == lisette.OptionSome {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:86:24
		child := subject.SomeVal
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:87:13
		next := newSnode[V]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:88:13
		next.Kids = slices.Clone(child.Kids)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:89:13
		next.Bucket = child.Bucket
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:90:13
		insertAt(&next, k, v, h, shift+5)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:91:13
		node.Kids[idx] = lisette.MakeOptionSome(&next)
	} else {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:94:13
		leaf := newSnode[V]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:95:13
		leaf.Bucket = append(leaf.Bucket, lisette.MakeTuple2[string, V](k, v))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:96:13
		node.Kids[idx] = lisette.MakeOptionSome(&leaf)
	}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:101:8
func Scount[V any](m SMap[V]) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:101:39
	return m.N
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:106:4
func visitKv[V any, R any](node *SNode[V], acc R, f func(R, string, V) R) R {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:107:5
	out := acc
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:108:5
	for _, item := range node.Bucket {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:108:5
		k := item.First
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:108:5
		v := item.Second
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:108:35
		out = f(out, k, v)
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:109:5
	for _, child := range node.Kids {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:110:9
		if child.Tag == lisette.OptionSome {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:110:34
			out = visitKv(child.SomeVal, out, f)
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:112:5
	return out
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:115:8
func ReduceKv[V any, R any](m SMap[V], zero R, f func(R, string, V) R) R {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:116:5
	return lisette.OptionMapOr(m.Root, zero, func(r *SNode[V]) R {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:116:29
		return visitKv(r, zero, f)
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:116:49*/)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:121:8
func FromEntries[V any](entries []lisette.Tuple2[string, V]) SMap[V] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:122:5
	m := EmptySmap[V]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:123:5
	for _, e := range entries {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:123:24
		m = Sassoc(m, e.First, e.Second)
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:124:5
	return m
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:130:8
func TransduceKeys[V any](m SMap[V], stages []Stage[string], zero int, f func(int, string) int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:131:5
	return ReduceKv(m, zero, func(acc int, k string, _ V) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:132:9
		st := FreshState()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:133:9
		value := lisette.MakeOptionSome(k)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:134:9
		for _, stage := range stages {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:135:13
			if value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:135:34
				break
			}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:136:13
			tmp := stage.Apply(st, valueOf(value))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:136:13
			next := tmp.First
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:136:13
			emitted := tmp.Second
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:136:13
			flag := tmp.Third
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:137:13
			st = next
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:138:13
			value = emitted
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:139:13
			if flag == MakeReducedStop() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:139:39
				break
			}
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:141:9
		if !value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:141:31
			return f(acc, valueOf(value))
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/smap.lis:141:64
		return acc
	})
}
