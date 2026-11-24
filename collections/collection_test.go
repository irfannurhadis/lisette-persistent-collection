package collections

import (
	"testing"

	lisette "github.com/ivov/lisette/prelude"
	"github.com/ivov/lisette/prelude/testkit"
)

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:5:4
func totalOf[A any](c Collection[A], zero A, f func(A, A) A) A {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:6:5
	return c.Reduce(zero, f)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:10:4
func consListIsPersistent(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:11:5
	a := Cons(1, Cons(2, EmptyList[int]()))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:12:5
	b := Cons(0, a)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:13:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:13:5
		assertLeft := lisette.OptionFromCommaOk[int](First(a))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:13:5
		assertRight := lisette.MakeOptionSome(1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:13:5
		if assertLeft != assertRight {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:13:5
			ctx.FailAssert(2, 441, 460, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:14:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:14:5
		assertLeft2 := lisette.OptionFromCommaOk[int](First(b))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:14:5
		assertRight2 := lisette.MakeOptionSome(0)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:14:5
		if assertLeft2 != assertRight2 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:14:5
			ctx.FailAssert(2, 472, 491, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft2)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight2)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:15:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:15:5
		assertLeft3 := ListCount(a)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:15:5
		if assertLeft3 != 2 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:15:5
			ctx.FailAssert(2, 503, 521, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft3)}, testkit.Operand{Label: "right", Value: lisette.Debug(2)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:16:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:16:5
		assertLeft4 := ListCount(b)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:16:5
		if assertLeft4 != 3 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:16:5
			ctx.FailAssert(2, 533, 551, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft4)}, testkit.Operand{Label: "right", Value: lisette.Debug(3)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:17:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:17:5
		assertLeft5 := totalOf(a, 0, func(x, y int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:17:44
			return x + y
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:17:49*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:17:5
		if assertLeft5 != 3 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:17:5
			ctx.FailAssert(2, 563, 606, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft5)}, testkit.Operand{Label: "right", Value: lisette.Debug(3)})
		}
	}
}

func TestConsListIsPersistent(t *testing.T) {
	defer testkit.Recover(t, 2, 336, 359)
	consListIsPersistent(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:21:4
func vectorTrieBasics(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:22:5
	v := EmptyVector[int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:23:5
	for i := range 40 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:23:22
		v = Vconj(v, i)
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:24:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:24:5
		assertLeft := lisette.OptionFromCommaOk[int](Vnth(v, 0))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:24:5
		assertRight := lisette.MakeOptionSome(0)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:24:5
		if assertLeft != assertRight {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:24:5
			ctx.FailAssert(2, 730, 751, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:25:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:25:5
		assertLeft2 := lisette.OptionFromCommaOk[int](Vnth(v, 31))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:25:5
		assertRight2 := lisette.MakeOptionSome(31)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:25:5
		if assertLeft2 != assertRight2 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:25:5
			ctx.FailAssert(2, 763, 786, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft2)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight2)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:26:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:26:5
		assertLeft3 := lisette.OptionFromCommaOk[int](Vnth(v, 32))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:26:5
		assertRight3 := lisette.MakeOptionSome(32)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:26:5
		if assertLeft3 != assertRight3 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:26:5
			ctx.FailAssert(2, 798, 821, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft3)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight3)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:27:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:27:5
		assertLeft4 := lisette.OptionFromCommaOk[int](Vnth(v, 39))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:27:5
		assertRight4 := lisette.MakeOptionSome(39)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:27:5
		if assertLeft4 != assertRight4 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:27:5
			ctx.FailAssert(2, 833, 856, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft4)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight4)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:28:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:28:5
		assertLeft5 := lisette.OptionFromCommaOk[int](Vnth(v, 40))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:28:5
		assertRight5 := lisette.MakeOptionNone[int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:28:5
		if assertLeft5 != assertRight5 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:28:5
			ctx.FailAssert(2, 868, 887, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft5)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight5)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:29:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:29:5
		assertLeft6 := totalOf(v, 0, func(x, y int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:29:44
			return x + y
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:29:49*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:29:5
		if assertLeft6 != 780 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:29:5
			ctx.FailAssert(2, 899, 944, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft6)}, testkit.Operand{Label: "right", Value: lisette.Debug(780)})
		}
	}
}

func TestVectorTrieBasics(t *testing.T) {
	defer testkit.Recover(t, 2, 621, 639)
	vectorTrieBasics(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:33:4
func hamtMapBasics(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:34:5
	m := EmptyMap[string, int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:35:5
	m = Massoc(m, "a", 1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:36:5
	m = Massoc(m, "b", 2)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:37:5
	m = Massoc(m, "a", 9)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:38:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:38:5
		assertLeft := lisette.OptionFromCommaOk[int](Mget(m, "a"))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:38:5
		assertRight := lisette.MakeOptionSome(9)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:38:5
		if assertLeft != assertRight {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:38:5
			ctx.FailAssert(2, 1109, 1132, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:39:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:39:5
		assertLeft2 := lisette.OptionFromCommaOk[int](Mget(m, "b"))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:39:5
		assertRight2 := lisette.MakeOptionSome(2)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:39:5
		if assertLeft2 != assertRight2 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:39:5
			ctx.FailAssert(2, 1144, 1167, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft2)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight2)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:40:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:40:5
		assertLeft3 := lisette.OptionFromCommaOk[int](Mget(m, "c"))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:40:5
		assertRight3 := lisette.MakeOptionNone[int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:40:5
		if assertLeft3 != assertRight3 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:40:5
			ctx.FailAssert(2, 1179, 1199, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft3)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight3)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:41:5
	m2 := Mdissoc(m, "a")
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:42:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:42:5
		assertLeft4 := lisette.OptionFromCommaOk[int](Mget(m2, "a"))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:42:5
		assertRight4 := lisette.MakeOptionNone[int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:42:5
		if assertLeft4 != assertRight4 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:42:5
			ctx.FailAssert(2, 1240, 1261, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft4)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight4)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:43:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:43:5
		assertLeft5 := lisette.OptionFromCommaOk[int](Mget(m, "a"))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:43:5
		assertRight5 := lisette.MakeOptionSome(9)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:43:5
		if assertLeft5 != assertRight5 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:43:5
			ctx.FailAssert(2, 1273, 1296, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft5)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight5)})
		}
	}
}

func TestHamtMapBasics(t *testing.T) {
	defer testkit.Recover(t, 2, 959, 974)
	hamtMapBasics(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:47:4
func setWrapsMap(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:48:5
	s := EmptySet[int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:49:5
	s = Sconj(s, 1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:50:5
	s = Sconj(s, 2)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:51:5
	if !Scontains(s, 1) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:51:5
		ctx.FailAssert(2, 1449, 1464, "bare", "assertion failed")
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:52:5
	if !Scontains(s, 2) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:52:5
		ctx.FailAssert(2, 1476, 1491, "bare", "assertion failed")
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:53:5
	if Scontains(s, 3) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:53:5
		ctx.FailAssert(2, 1503, 1519, "bare", "assertion failed")
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:54:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:54:5
		assertLeft := s.Backing.N
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:54:5
		if assertLeft != 2 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:54:5
			ctx.FailAssert(2, 1531, 1547, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft)}, testkit.Operand{Label: "right", Value: lisette.Debug(2)})
		}
	}
}

func TestSetWrapsMap(t *testing.T) {
	defer testkit.Recover(t, 2, 1347, 1360)
	setWrapsMap(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:58:4
func queueTwoList(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:59:5
	q := EmptyQueue[int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:60:5
	q = Qconj(q, 1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:61:5
	q = Qconj(q, 2)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:62:5
	q = Qconj(q, 3)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:63:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:63:5
		assertLeft := lisette.OptionFromCommaOk[int](Qpeek(q))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:63:5
		assertRight := lisette.MakeOptionSome(1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:63:5
		if assertLeft != assertRight {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:63:5
			ctx.FailAssert(2, 1687, 1706, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:64:5
	q = Qpop(q)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:65:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:65:5
		assertLeft2 := lisette.OptionFromCommaOk[int](Qpeek(q))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:65:5
		assertRight2 := lisette.MakeOptionSome(2)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:65:5
		if assertLeft2 != assertRight2 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:65:5
			ctx.FailAssert(2, 1734, 1753, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft2)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight2)})
		}
	}
}

func TestQueueTwoList(t *testing.T) {
	defer testkit.Recover(t, 2, 1562, 1576)
	queueTwoList(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:69:4
func oneTransducerOverFiveCollections(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:71:5
	stages := []Stage[int]{{Apply: /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:72:16*/ func(st RState, x int) lisette.Tuple3[RState, lisette.Option[int], Reduced] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:72:37
		return lisette.MakeTuple3[RState, lisette.Option[int], Reduced](st, lisette.MakeOptionSome(x*2), MakeReducedContinue())
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:72:72*/}}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:74:5
	v := EmptyVector[int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:75:5
	for i := range 5 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:75:21
		v = Vconj(v, i)
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:76:5
	l := Cons(0, Cons(1, Cons(2, Cons(3, Cons(4, EmptyList[int]())))))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:77:5
	m := EmptyMap[string, int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:78:5
	m = Massoc(m, "a", 1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:79:5
	stringStages := []Stage[string]{{Apply: /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:80:16*/ func(st RState, k string) lisette.Tuple3[RState, lisette.Option[string], Reduced] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:80:40
		return lisette.MakeTuple3[RState, lisette.Option[string], Reduced](st, lisette.MakeOptionSome(k), MakeReducedContinue())
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:80:71*/}}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:82:5
	s := EmptySet[int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:83:5
	s = Sconj(s, 1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:85:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:85:5
		assertLeft := TransduceVec(v, stages, 0, func(acc, x int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:85:59
			return acc + x
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:85:66*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:85:5
		if assertLeft != 20 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:85:5
			ctx.FailAssert(2, 2378, 2439, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft)}, testkit.Operand{Label: "right", Value: lisette.Debug(20)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:86:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:86:5
		assertLeft2 := TransduceList(l, stages, 0, func(acc, x int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:86:60
			return acc + x
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:86:67*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:86:5
		if assertLeft2 != 20 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:86:5
			ctx.FailAssert(2, 2451, 2513, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft2)}, testkit.Operand{Label: "right", Value: lisette.Debug(20)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:87:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:87:5
		assertLeft3 := TransduceMapKeys(m, stringStages, 0, func(acc int, _ string) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:87:75
			return acc + 1
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:87:82*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:87:5
		if assertLeft3 != 1 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:87:5
			ctx.FailAssert(2, 2525, 2601, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft3)}, testkit.Operand{Label: "right", Value: lisette.Debug(1)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:89:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:89:5
		assertLeft4 := m.ReduceKv(0, func(acc int, _ string, val int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:89:60
			return acc + val
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:89:69*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:89:5
		if assertLeft4 != 1 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:89:5
			ctx.FailAssert(2, 2692, 2755, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft4)}, testkit.Operand{Label: "right", Value: lisette.Debug(1)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:90:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:90:5
		assertLeft5 := TransduceSet(s, stages, 0, func(acc, k int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:90:59
			return acc + k
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:90:66*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:90:5
		if assertLeft5 != 2 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:90:5
			ctx.FailAssert(2, 2767, 2827, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft5)}, testkit.Operand{Label: "right", Value: lisette.Debug(2)})
		}
	}
}

func TestOneTransducerOverFiveCollections(t *testing.T) {
	defer testkit.Recover(t, 2, 1768, 1804)
	oneTransducerOverFiveCollections(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:94:4
func intoMaterializesPipeline(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:95:5
	stages := []Stage[int]{{Apply: /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:96:16*/ func(st RState, x int) lisette.Tuple3[RState, lisette.Option[int], Reduced] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:97:13
		if x > 2 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:97:24
			return lisette.MakeTuple3[RState, lisette.Option[int], Reduced](st, lisette.MakeOptionNone[int](), MakeReducedContinue())
		}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:97:62
		return lisette.MakeTuple3[RState, lisette.Option[int], Reduced](st, lisette.MakeOptionSome(x+10), MakeReducedContinue())
	}}}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:100:5
	out := Into(stages, []int{0, 1, 2, 3, 4})
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:101:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:101:5
		assertLeft := lisette.OptionFromCommaOk[int](Vnth(out, 0))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:101:5
		assertRight := lisette.MakeOptionSome(10)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:101:5
		if assertLeft != assertRight {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:101:5
			ctx.FailAssert(2, 3110, 3134, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:102:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:102:5
		assertLeft2 := lisette.OptionFromCommaOk[int](Vnth(out, 1))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:102:5
		assertRight2 := lisette.MakeOptionSome(11)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:102:5
		if assertLeft2 != assertRight2 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:102:5
			ctx.FailAssert(2, 3146, 3170, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft2)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight2)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:103:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:103:5
		assertLeft3 := lisette.OptionFromCommaOk[int](Vnth(out, 2))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:103:5
		assertRight3 := lisette.MakeOptionSome(12)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:103:5
		if assertLeft3 != assertRight3 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:103:5
			ctx.FailAssert(2, 3182, 3206, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft3)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight3)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:104:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:104:5
		assertLeft4 := lisette.OptionFromCommaOk[int](Vnth(out, 3))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:104:5
		assertRight4 := lisette.MakeOptionNone[int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:104:5
		if assertLeft4 != assertRight4 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/collection.test.lis:104:5
			ctx.FailAssert(2, 3218, 3238, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft4)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight4)})
		}
	}
}

func TestIntoMaterializesPipeline(t *testing.T) {
	defer testkit.Recover(t, 2, 2842, 2868)
	intoMaterializesPipeline(testkit.New(t))
}
