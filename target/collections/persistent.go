package collections

import (
	"fmt"
	"slices"

	lisette "github.com/ivov/lisette/prelude"
)

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:13:5
func MakeListEmpty[T any]() List[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:13:5
	return List[T]{Tag: ListEmpty}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:14:5
func MakeListCons[T any](arg0 T, arg1 *List[T]) List[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:14:5
	return List[T]{Tag: ListCons, Head: arg0, Tail: arg1}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:238:31
func MakeUnitValOnly() UnitVal {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:238:31
	return UnitVal{Tag: UnitValOnly}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
type ListTag uint8

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
const (
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
	ListEmpty ListTag = iota
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
	ListCons
)

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
type List[T any] struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
	Tag ListTag
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
	Head T
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
	Tail *List[T]
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
func (l List[T]) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
	switch l.Tag {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
	case ListEmpty:
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
		return "Empty"
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
	case ListCons:
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
		return fmt.Sprintf("Cons { head: %s, tail: %s }", lisette.Debug(l.Head), lisette.Debug(l.Tail))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
	default:
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:11:10
		return fmt.Sprintf("List(%d)", l.Tag)
	}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:17:8
func EmptyList[T any]() List[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:17:37
	return MakeListEmpty[T]()
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:19:8
func Cons[T any](head T, tail List[T]) List[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:20:5
	return List[T]{
		Tag: ListCons,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:20:17
		Head: head,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:20:23
		Tail: &tail,
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:20:36*/
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:23:8
func First[T any](xs List[T]) (T, bool) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:24:5
	if xs.Tag == ListEmpty {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:25:23
		var zero T
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:25:23
		return zero, false
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:26:35
	return xs.Head, true
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:30:8
func Rest[T any](xs List[T]) List[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:31:5
	if xs.Tag == ListEmpty {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:32:23
		return MakeListEmpty[T]()
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:33:35
	return *xs.Tail
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:37:8
func ListCount[T any](xs List[T]) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:38:5
	if xs.Tag == ListEmpty {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:39:23
		return 0
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:40:35
	return 1 + ListCount(*xs.Tail)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:44:8
func FromSlice[T any](xs []T) List[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:45:5
	out := EmptyList[T]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:46:5
	i := len(xs) - 1
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:47:5
	for i >= 0 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:48:9
		out = Cons(xs[i], out)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:49:9
		i--
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:51:5
	return out
}

const Bits int = 5

const Width int = 32

const Mask int = 31

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:62:12
type VNode[T any] struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:62:12
	Children []lisette.Option[*VNode[T]]
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:62:12
	Leaf []lisette.Option[T]
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:62:12
func (v VNode[T]) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:62:12
	return fmt.Sprintf("VNode { children: %s, leaf: %s }", lisette.Debug(v.Children), lisette.Debug(v.Leaf))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:68:12
type Vector[T any] struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:68:12
	N int
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:68:12
	Shift int
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:68:12
	Root lisette.Option[*VNode[T]]
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:68:12
	Tail []T
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:68:12
func (v Vector[T]) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:68:12
	return fmt.Sprintf("Vector { n: %s, shift: %s, root: %s, tail: %s }", lisette.Debug(v.N), lisette.Debug(v.Shift), lisette.Debug(v.Root), lisette.Debug(v.Tail))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:75:8
func EmptyVector[T any]() Vector[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:76:5
	return Vector[T]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:76:14
		N: 0,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:76:20
		Shift: Bits,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:76:45
		Tail: []T{},
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:76:67*/
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:79:4
func newNode[T any]() VNode[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:80:5
	children := []lisette.Option[*VNode[T]]{}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:81:5
	for range Width {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:81:25
		children = append(children, lisette.MakeOptionNone[*VNode[T]]())
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:82:5
	return VNode[T]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:82:13
		Children: children,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:82:23
		Leaf: []lisette.Option[T]{},
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:82:53*/
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:86:4
func pushTail[T any](node lisette.Option[*VNode[T]], level int, tail []T, count int) *VNode[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:87:5
	out := newNode[T]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:88:5
	if node.Tag == lisette.OptionSome {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:89:20
		r := node.SomeVal
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:90:13
		out.Children = slices.Clone(r.Children)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:91:13
		out.Leaf = r.Leaf
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:95:5
	idx := (count - 1) >> level & Mask
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:96:5
	if level == Bits {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:98:9
		leafNode := newNode[T]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:99:9
		leaf := []lisette.Option[T]{}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:100:9
		for _, x := range tail {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:100:25
			leaf = append(leaf, lisette.MakeOptionSome(x))
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:101:9
		for len(leaf) < Width {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:101:39
			leaf = append(leaf, lisette.MakeOptionNone[T]())
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:102:9
		leafNode.Leaf = leaf
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:103:9
		out.Children[idx] = lisette.MakeOptionSome(&leafNode)
	} else {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:105:9
		child := out.Children[idx]
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:106:9
		base := out.Children
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:106:9
		base[idx] = lisette.MakeOptionSome(pushTail(child, level-Bits, tail, count))
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:108:5
	return &out
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:111:8
func Vconj[T any](v Vector[T], x T) Vector[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:112:5
	if len(v.Tail) < Width {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:113:9
		tail := append(slices.Clip(v.Tail), x)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:114:9
		return Vector[T]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:114:25
			N: v.N + 1,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:114:37
			Shift: v.Shift,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:114:53
			Root: v.Root,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:114:67
			Tail: tail,
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:114:73*/
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:117:5
	full := v.Tail
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:118:5
	var newRoot lisette.Option[*VNode[T]]
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:118:20
	if v.Root.Tag == lisette.OptionNone {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:120:13
		leafNode := newNode[T]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:121:13
		leaf := []lisette.Option[T]{}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:122:13
		for _, y := range full {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:122:29
			leaf = append(leaf, lisette.MakeOptionSome(y))
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:123:13
		for len(leaf) < Width {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:123:43
			leaf = append(leaf, lisette.MakeOptionNone[T]())
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:124:13
		leafNode.Leaf = leaf
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:125:13
		rootNode := newNode[T]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:126:13
		rootNode.Children[0] = lisette.MakeOptionSome(&leafNode)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:127:13
		newRoot = lisette.MakeOptionSome(&rootNode)
	} else {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:129:20
		newRoot = lisette.MakeOptionSome[*VNode[T]](pushTail(lisette.MakeOptionSome(v.Root.SomeVal), v.Shift, full, v.N+1))
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:131:5
	return Vector[T]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:131:14
		N: v.N + 1,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:131:26
		Shift: v.Shift,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:131:42
		Root: newRoot,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:131:58
		Tail: []T{x},
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:131:69*/
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:134:8
func VtoSlice[T any](v Vector[T]) []T {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:135:5
	out := []T{}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:136:5
	i := 0
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:137:5
	for i < v.N {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:138:9
		if x, ok := Vnth(v, i); ok {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:138:39
			out = append(out, x)
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:139:9
		i++
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:141:5
	return out
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:144:8
func Vnth[T any](v Vector[T], i int) (T, bool) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:145:5
	if i < 0 || i >= v.N {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:145:28
		var zero2 T
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:145:28
		return zero2, false
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:146:5
	tailOff := v.N - len(v.Tail)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:147:5
	if i >= tailOff {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:147:24
		v2 := lisette.SliceGet(v.Tail, i-tailOff)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:147:24
		if v2.Tag == lisette.OptionSome {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:147:24
			return v2.SomeVal, true
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:147:24
		var zero3 T
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:147:24
		return zero3, false
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:148:5
	check := v.Root
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:148:5
	if check.Tag != lisette.OptionSome {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:148:5
		var zero4 T
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:148:5
		return zero4, false
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:148:5
	node := check.SomeVal
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:149:5
	level := v.Shift
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:150:5
	for level > Bits {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:151:9
		idx := i >> level & Mask
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:152:9
		check2 := node.Children[idx]
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:152:9
		if check2.Tag != lisette.OptionSome {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:152:9
			var zero5 T
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:152:9
			return zero5, false
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:152:9
		node = check2.SomeVal
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:153:9
		level -= Bits
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:156:5
	check3 := node.Children[i>>Bits&Mask]
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:156:5
	if check3.Tag != lisette.OptionSome {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:156:5
		var zero6 T
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:156:5
		return zero6, false
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:156:5
	holder := *check3.SomeVal
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:157:5
	v3 := holder.Leaf[i&Mask]
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:157:5
	if v3.Tag == lisette.OptionSome {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:157:5
		return v3.SomeVal, true
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:157:5
	var zero T
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:157:5
	return zero, false
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:165:12
type Entry[K any, V any] struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:165:12
	Key K
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:165:12
	Value V
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:165:12
func (e Entry[K, V]) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:165:12
	return fmt.Sprintf("Entry { key: %s, value: %s }", lisette.Debug(e.Key), lisette.Debug(e.Value))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:167:12
type HNode[K any, V any] struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:167:12
	Children []lisette.Option[*HNode[K, V]]
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:167:12
	Buckets []Entry[K, V]
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:167:12
func (h HNode[K, V]) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:167:12
	return fmt.Sprintf("HNode { children: %s, buckets: %s }", lisette.Debug(h.Children), lisette.Debug(h.Buckets))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:172:12
type HashMap[K comparable, V any] struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:172:12
	N int
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:172:12
	Compact []Entry[K, V]
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:172:12
	Root lisette.Option[*HNode[K, V]]
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:172:12
func (h HashMap[K, V]) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:172:12
	return fmt.Sprintf("HashMap { n: %s, compact: %s, root: %s }", lisette.Debug(h.N), lisette.Debug(h.Compact), lisette.Debug(h.Root))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:178:8
func EmptyMap[K comparable, V any]() HashMap[K, V] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:179:5
	return HashMap[K, V]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:179:15
		N: 0,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:179:21
		Compact: []Entry[K, V]{},
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:179:68*/
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:182:8
func Mget[K comparable, V any](m HashMap[K, V], k K) (V, bool) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:183:5
	for _, e := range m.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:184:9
		if e.Key == k {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:184:25
			return e.Value, true
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:186:5
	check := m.Root
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:186:5
	if check.Tag != lisette.OptionSome {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:186:5
		var zero V
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:186:5
		return zero, false
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:186:5
	node := check.SomeVal
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:187:5
	for {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:188:9
		for _, e := range node.Buckets {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:189:13
			if e.Key == k {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:189:29
				return e.Value, true
			}
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:192:9
		next := lisette.MakeOptionNone[*HNode[K, V]]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:193:9
		for _, c := range node.Children {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:194:13
			if c.IsSome() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:194:30
				next = c
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:194:40
				break
			}
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:196:9
		if next.Tag == lisette.OptionNone {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:196:51
			var zero2 V
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:196:51
			return zero2, false
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:196:33
		node = next.SomeVal
	}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:200:8
func Massoc[K comparable, V any](m HashMap[K, V], k K, v V) HashMap[K, V] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:201:5
	if m.N <= 8 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:202:9
		compact := []Entry[K, V]{}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:203:9
		found := false
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:204:9
		for _, e := range m.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:205:13
			if e.Key == k {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:205:29
				compact = append(compact, Entry[K, V]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:205:62
					Key: k,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:205:70
					Value: v,
				} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:205:80*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:205:83
				found = true
			} else {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:206:20
				compact = append(compact, e)
			}
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:208:9
		if found {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:208:20
			return HashMap[K, V]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:208:37
				N: m.N,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:208:45
				Compact: compact,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:208:54
				Root: m.Root,
			} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:208:68*/
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:209:9
		if m.N < 8 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:209:22
			compact = append(compact, Entry[K, V]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:209:55
				Key: k,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:209:63
				Value: v,
			} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:209:73*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:209:76
			return HashMap[K, V]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:209:93
				N: m.N + 1,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:209:105
				Compact: compact,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:209:114
				Root: m.Root,
			} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:209:128*/
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:212:5
	node := HNode[K, V]{
		Children: []lisette.Option[*HNode[K, V]]{},
		Buckets:  []Entry[K, V]{},
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:216:5
	if m.Root.Tag == lisette.OptionSome {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:217:22
		node.Buckets = m.Root.SomeVal.Buckets
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:220:5
	node.Buckets = append(node.Buckets, Entry[K, V]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:220:48
		Key: k,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:220:56
		Value: v,
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:220:66*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:221:5
	return HashMap[K, V]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:221:15
		N: m.N + 1,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:221:27
		Compact: m.Compact,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:221:47
		Root: lisette.MakeOptionSome(&node),
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:221:66*/
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:224:8
func Mdissoc[K comparable, V any](m HashMap[K, V], k K) HashMap[K, V] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:225:5
	compact := []Entry[K, V]{}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:226:5
	removed := false
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:227:5
	for _, e := range m.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:228:9
		if e.Key == k {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:228:25
			removed = true
		} else {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:228:49
			compact = append(compact, e)
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:230:5
	if removed {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:230:18
		return HashMap[K, V]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:230:35
			N: m.N - 1,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:230:47
			Compact: compact,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:230:56
			Root: m.Root,
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:230:70*/
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:231:5
	return m
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:238:10
type UnitValTag uint8

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:238:10
const (
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:238:10
	UnitValOnly UnitValTag = iota
)

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:238:10
type UnitVal struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:238:10
	Tag UnitValTag
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:238:10
func (u UnitVal) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:238:10
	switch u.Tag {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:238:10
	case UnitValOnly:
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:238:10
		return "Only"
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:238:10
	default:
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:238:10
		return fmt.Sprintf("UnitVal(%d)", u.Tag)
	}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:240:12
type HashSet[T comparable] struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:240:12
	Backing HashMap[T, UnitVal]
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:240:12
func (h HashSet[T]) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:240:12
	return fmt.Sprintf("HashSet { backing: %s }", lisette.Debug(h.Backing))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:242:8
func EmptySet[T comparable]() HashSet[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:243:5
	return HashSet[T]{Backing: EmptyMap[T, UnitVal]()}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:246:8
func Sconj[T comparable](s HashSet[T], x T) HashSet[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:247:5
	return HashSet[T]{Backing: Massoc(s.Backing, x, MakeUnitValOnly())}
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:250:8
func Scontains[T comparable](s HashSet[T], x T) bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:251:5
	_, ok := Mget(s.Backing, x)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:251:5
	return ok
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:258:12
type Queue[T any] struct {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:258:12
	Front List[T]
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:258:12
	Rear List[T]
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:258:12
func (q Queue[T]) DebugString() string {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:258:12
	return fmt.Sprintf("Queue { front: %s, rear: %s }", lisette.Debug(q.Front), lisette.Debug(q.Rear))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:260:8
func EmptyQueue[T any]() Queue[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:261:5
	return Queue[T]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:261:13
		Front: EmptyList[T](),
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:261:37
		Rear: EmptyList[T](),
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:261:60*/
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:264:4
func normalize[T any](q Queue[T]) Queue[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:265:5
	if _, ok := First(q.Front); ok {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:266:20
		return q
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:267:17
	return Queue[T]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:267:25
		Front: FromSlice(reverseSlice(toSlice(q.Rear))),
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:267:77
		Rear: EmptyList[T](),
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:267:100*/
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:271:8
func Qconj[T any](q Queue[T], x T) Queue[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:272:5
	return normalize(Queue[T]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:272:23
		Front: q.Front,
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:272:39
		Rear: Cons(x, q.Rear),
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:272:62*/)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:275:8
func Qpeek[T any](q Queue[T]) (T, bool) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:275:45
	return First(q.Front)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:277:8
func Qpop[T any](q Queue[T]) Queue[T] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:277:43
	return normalize(Queue[T]{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:277:61
		Front: Rest(q.Front),
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:277:83
		Rear: q.Rear,
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:277:97*/)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:279:4
func reverseSlice[T any](xs []T) []T {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:280:5
	out := []T{}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:281:5
	i := len(xs) - 1
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:282:5
	for i >= 0 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:282:20
		out = append(out, xs[i])
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:282:45
		i--
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:283:5
	return out
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:286:4
func toSlice[T any](xs List[T]) []T {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:287:5
	out := []T{}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:288:5
	cur := xs
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:289:5
	for {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:290:9
		if cur.Tag == ListEmpty {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:291:27
			return out
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:292:41
		tail := cur.Tail
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:292:43
		out = append(out, cur.Head)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/persistent.lis:292:67
		cur = *tail
	}
}
