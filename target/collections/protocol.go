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

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:33:8
func (l List[T]) Count() int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:33:29
	return ListCount(l)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:34:8
func (l List[T]) Empty() bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:34:30
	if l.Tag == ListEmpty {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:34:57
		return true
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:34:68
	return false
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:35:8
func (l List[T]) Reduce(zero T, f func(T, T) T) T {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:36:9
	acc := zero
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:37:9
	cur := l
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:38:9
	for {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:39:13
		if cur.Tag == ListEmpty {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:40:31
			break
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:41:45
		tail := cur.Tail
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:41:47
		acc = f(acc, cur.Head)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:41:67
		cur = *tail
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:44:9
	return acc
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:49:8
func (v Vector[T]) Count() int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:49:29
	return v.N
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:50:8
func (v Vector[T]) Empty() bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:50:30
	return v.N == 0
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:51:8
func (v Vector[T]) Reduce(zero T, f func(T, T) T) T {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:53:9
	acc := zero
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:54:9
	i := 0
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:55:9
	for i < v.N {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:56:13
		if x, ok := Vnth(v, i); ok {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:56:46
			acc = f(acc, x)
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:57:13
		i++
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:59:9
	return acc
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:64:8
func (h HashMap[K, V]) Count() int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:64:29
	return h.N
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:65:8
func (h HashMap[K, V]) Empty() bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:65:30
	return h.N == 0
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:66:8
func (h HashMap[K, V]) ReduceKv(zero V, f func(V, K, V) V) V {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:67:9
	acc := zero
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:68:9
	for _, e := range h.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:68:33
		acc = f(acc, e.Key, e.Value)
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:69:9
	return acc
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:72:8
func (h HashMap[K, V]) Reduce(zero V, f func(V, V) V) V {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:73:9
	acc := zero
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:74:9
	for _, e := range h.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:74:33
		acc = f(acc, e.Value)
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:75:9
	return acc
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:80:8
func (h HashSet[K]) Count() int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:80:29
	return h.Backing.N
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:81:8
func (h HashSet[K]) Empty() bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:81:30
	return h.Backing.N == 0
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:82:8
func (h HashSet[K]) Reduce(zero K, f func(K, K) K) K {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:83:9
	acc := zero
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:84:9
	for _, e := range h.Backing.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:84:41
		acc = f(acc, e.Key)
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:85:9
	return acc
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:90:8
func (q Queue[T]) Count() int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:90:29
	return ListCount(q.Front) + ListCount(q.Rear)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:91:8
func (q Queue[T]) Empty() bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:91:30
	_, ok := First(q.Front)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:91:30
	return !ok
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:92:8
func (q Queue[T]) Reduce(zero T, f func(T, T) T) T {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:93:9
	acc := zero
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:94:9
	cur := q.Front
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:95:9
	for {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:96:13
		if cur.Tag == ListEmpty {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:97:31
			break
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:98:45
		tail := cur.Tail
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:98:47
		acc = f(acc, cur.Head)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:98:67
		cur = *tail
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:101:9
	return acc
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:110:8
func TransduceVec[A any, R any](v Vector[A], stages []Stage[A], zero R, f func(R, A) R) R {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:111:5
	acc := zero
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:112:5
	st := FreshState()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:113:5
	i := 0
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:114:5
	for i < v.N {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:115:9
		if st.Stopped {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:115:25
			break
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:116:9
		value := lisette.OptionFromCommaOk[A](Vnth(v, i))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:117:9
		for _, stage := range stages {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:118:13
			if value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:118:34
				break
			}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:119:13
			tmp := stage.Apply(st, valueOf(value))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:119:13
			next := tmp.First
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:119:13
			emitted := tmp.Second
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:119:13
			flag := tmp.Third
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:120:13
			st = next
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:121:13
			value = emitted
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:122:13
			if flag == MakeReducedStop() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:122:39
				copy2 := st
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:122:39
				copy2.Stopped = true
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:122:39
				st = copy2
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:122:76
				break
			}
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:124:9
		if !value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:124:31
			acc = f(acc, valueOf(value))
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:125:9
		i++
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:127:5
	return acc
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:130:8
func TransduceList[A any, R any](xs List[A], stages []Stage[A], zero R, f func(R, A) R) R {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:131:5
	acc := zero
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:132:5
	st := FreshState()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:133:5
	cur := xs
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:134:5
	for {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:135:9
		if cur.Tag == ListEmpty {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:136:27
			break
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:137:41
		tail := cur.Tail
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:138:17
		if st.Stopped {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:138:33
			break
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:139:17
		value := lisette.MakeOptionSome(cur.Head)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:140:17
		for _, stage := range stages {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:141:21
			if value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:141:42
				break
			}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:142:21
			tmp := stage.Apply(st, valueOf(value))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:142:21
			next := tmp.First
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:142:21
			emitted := tmp.Second
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:142:21
			flag := tmp.Third
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:143:21
			st = next
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:144:21
			value = emitted
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:145:21
			if flag == MakeReducedStop() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:145:47
				copy2 := st
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:145:47
				copy2.Stopped = true
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:145:47
				st = copy2
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:145:84
				break
			}
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:147:17
		if !value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:147:39
			acc = f(acc, valueOf(value))
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:148:17
		cur = *tail
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:152:5
	return acc
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:155:8
func TransduceMapKeys[K comparable, R any](m HashMap[K, R], stages []Stage[K], zero R, f func(R, K) R) R {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:157:5
	acc := zero
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:158:5
	st := FreshState()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:159:5
	for _, e := range m.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:160:9
		if st.Stopped {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:160:25
			break
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:161:9
		value := lisette.MakeOptionSome(e.Key)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:162:9
		for _, stage := range stages {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:163:13
			if value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:163:34
				break
			}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:164:13
			tmp := stage.Apply(st, valueOf(value))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:164:13
			next := tmp.First
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:164:13
			emitted := tmp.Second
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:164:13
			flag := tmp.Third
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:165:13
			st = next
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:166:13
			value = emitted
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:167:13
			if flag == MakeReducedStop() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:167:39
				copy2 := st
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:167:39
				copy2.Stopped = true
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:167:39
				st = copy2
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:167:76
				break
			}
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:169:9
		if !value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:169:31
			acc = f(acc, valueOf(value))
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:171:5
	return acc
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:174:8
func TransduceSet[K comparable, R any](s HashSet[K], stages []Stage[K], zero R, f func(R, K) R) R {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:176:5
	acc := zero
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:177:5
	st := FreshState()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:178:5
	for _, e := range s.Backing.Compact {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:179:9
		if st.Stopped {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:179:25
			break
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:180:9
		value := lisette.MakeOptionSome(e.Key)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:181:9
		for _, stage := range stages {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:182:13
			if value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:182:34
				break
			}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:183:13
			tmp := stage.Apply(st, valueOf(value))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:183:13
			next := tmp.First
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:183:13
			emitted := tmp.Second
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:183:13
			flag := tmp.Third
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:184:13
			st = next
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:185:13
			value = emitted
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:186:13
			if flag == MakeReducedStop() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:186:39
				copy2 := st
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:186:39
				copy2.Stopped = true
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:186:39
				st = copy2
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:186:76
				break
			}
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:188:9
		if !value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:188:31
			acc = f(acc, valueOf(value))
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:190:5
	return acc
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:194:8
func Into[A any](stages []Stage[A], xs []A) Vector[A] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:195:5
	v := EmptyVector[A]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:196:5
	st := FreshState()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:197:5
	for _, a := range xs {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:198:9
		if st.Stopped {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:198:25
			break
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:199:9
		value := lisette.MakeOptionSome(a)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:200:9
		for _, stage := range stages {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:201:13
			if value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:201:34
				break
			}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:202:13
			tmp := stage.Apply(st, valueOf(value))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:202:13
			next := tmp.First
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:202:13
			emitted := tmp.Second
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:202:13
			flag := tmp.Third
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:203:13
			st = next
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:204:13
			value = emitted
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:205:13
			if flag == MakeReducedStop() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:205:39
				copy2 := st
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:205:39
				copy2.Stopped = true
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:205:39
				st = copy2
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:205:76
				break
			}
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:207:9
		if !value.IsNone() {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:207:31
			v = Vconj(v, valueOf(value))
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/protocol.lis:209:5
	return v
}
