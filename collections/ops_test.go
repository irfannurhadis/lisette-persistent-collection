package collections

import (
	"strings"
	"testing"

	lisette "github.com/ivov/lisette/prelude"
	"github.com/ivov/lisette/prelude/testkit"
)

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:6:4
func setUnionIntersectionDifference(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:7:5
	a := EmptySet[int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:8:5
	a = Sconj(a, 1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:8:22
	a = Sconj(a, 2)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:9:5
	b := EmptySet[int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:10:5
	b = Sconj(b, 2)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:10:22
	b = Sconj(b, 3)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:11:5
	u := Sunion(a, b)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:12:5
	if !(Scontains(u, 1) && Scontains(u, 2) && Scontains(u, 3)) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:12:5
		ctx.FailAssert(4, 318, 371, "bare", "assertion failed")
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:13:5
	i := Sintersection(a, b)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:14:5
	if !(Scontains(i, 2) && !Scontains(i, 1) && !Scontains(i, 3)) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:14:5
		ctx.FailAssert(4, 415, 470, "bare", "assertion failed")
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:15:5
	d := Sdifference(a, b)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:16:5
	if !(Scontains(d, 1) && !Scontains(d, 2)) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:16:5
		ctx.FailAssert(4, 512, 547, "bare", "assertion failed")
	}
}

func TestSetUnionIntersectionDifference(t *testing.T) {
	defer testkit.Recover(t, 4, 104, 137)
	setUnionIntersectionDifference(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:20:4
func mapMergeIsFunctional(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:21:5
	a := EmptyMap[string, int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:22:5
	a = Massoc(a, "a", 1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:23:5
	b := EmptyMap[string, int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:24:5
	b = Massoc(b, "a", 9)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:24:28
	b = Massoc(b, "b", 2)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:25:5
	m := Mmerge(a, b)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:26:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:26:5
		assertLeft := lisette.OptionFromCommaOk[int](Mget(m, "a"))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:26:5
		assertRight := lisette.MakeOptionSome(9)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:26:5
		if assertLeft != assertRight {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:26:5
			ctx.FailAssert(4, 783, 806, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:27:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:27:5
		assertLeft2 := lisette.OptionFromCommaOk[int](Mget(m, "b"))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:27:5
		assertRight2 := lisette.MakeOptionSome(2)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:27:5
		if assertLeft2 != assertRight2 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:27:5
			ctx.FailAssert(4, 818, 841, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft2)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight2)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:28:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:28:5
		assertLeft3 := lisette.OptionFromCommaOk[int](Mget(a, "a"))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:28:5
		assertRight3 := lisette.MakeOptionSome(1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:28:5
		if assertLeft3 != assertRight3 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:28:5
			ctx.FailAssert(4, 853, 876, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft3)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight3)})
		}
	}
}

func TestMapMergeIsFunctional(t *testing.T) {
	defer testkit.Recover(t, 4, 562, 585)
	mapMergeIsFunctional(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:32:4
func mapKeyValueTransforms(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:33:5
	m := EmptyMap[string, int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:34:5
	m = Massoc(m, "a", 1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:34:28
	m = Massoc(m, "bb", 2)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:35:5
	upper := MapKeys(m, strings.ToUpper)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:36:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:36:5
		assertLeft := lisette.OptionFromCommaOk[int](Mget(upper, "A"))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:36:5
		assertRight := lisette.MakeOptionSome(1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:36:5
		if assertLeft != assertRight {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:36:5
			ctx.FailAssert(4, 1091, 1118, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:37:5
	doubled := MapValues(m, func(v int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:37:42
		return v * 2
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:37:47*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:38:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:38:5
		assertLeft2 := lisette.OptionFromCommaOk[int](Mget(doubled, "a"))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:38:5
		assertRight2 := lisette.MakeOptionSome(2)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:38:5
		if assertLeft2 != assertRight2 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:38:5
			ctx.FailAssert(4, 1178, 1207, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft2)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight2)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:39:5
	short := FilterKeys(m, func(k string) bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:39:44
		return len(k) == 1
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:39:59*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:40:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:40:5
		assertLeft3 := lisette.OptionFromCommaOk[int](Mget(short, "a"))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:40:5
		assertRight3 := lisette.MakeOptionSome(1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:40:5
		if assertLeft3 != assertRight3 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:40:5
			ctx.FailAssert(4, 1279, 1306, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft3)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight3)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:41:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:41:5
		assertLeft4 := lisette.OptionFromCommaOk[int](Mget(short, "bb"))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:41:5
		assertRight4 := lisette.MakeOptionNone[int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:41:5
		if assertLeft4 != assertRight4 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:41:5
			ctx.FailAssert(4, 1318, 1343, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft4)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight4)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:42:5
	big := FilterValues(m, func(v int) bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:42:41
		return v > 1
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:42:46*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:43:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:43:5
		assertLeft5 := lisette.OptionFromCommaOk[int](Mget(big, "bb"))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:43:5
		assertRight5 := lisette.MakeOptionSome(2)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:43:5
		if assertLeft5 != assertRight5 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:43:5
			ctx.FailAssert(4, 1402, 1428, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft5)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight5)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:44:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:44:5
		assertLeft6 := lisette.OptionFromCommaOk[int](Mget(big, "a"))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:44:5
		assertRight6 := lisette.MakeOptionNone[int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:44:5
		if assertLeft6 != assertRight6 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:44:5
			ctx.FailAssert(4, 1440, 1462, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft6)}, testkit.Operand{Label: "right", Value: lisette.Debug(assertRight6)})
		}
	}
}

func TestMapKeyValueTransforms(t *testing.T) {
	defer testkit.Recover(t, 4, 915, 939)
	mapKeyValueTransforms(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:48:4
func smapReduceKvAndTransduce(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:49:5
	m := EmptySmap[int]()
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:50:5
	m = Sassoc(m, "a", 1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:50:28
	m = Sassoc(m, "b", 2)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:51:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:51:5
		assertLeft := ReduceKv(m, 0, func(acc int, _ string, v int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:51:59
			return acc + v
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:51:66*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:51:5
		if assertLeft != 3 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:51:5
			ctx.FailAssert(4, 1604, 1664, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft)}, testkit.Operand{Label: "right", Value: lisette.Debug(3)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:52:5
	stages := []Stage[string]{{Apply: func(st RState, k string) lisette.Tuple3[RState, lisette.Option[string], Reduced] {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:52:58
		return lisette.MakeTuple3[RState, lisette.Option[string], Reduced](st, lisette.MakeOptionSome(k), MakeReducedContinue())
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:52:89*/}}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:53:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:53:5
		assertLeft2 := TransduceKeys(m, stages, 0, func(acc int, _ string) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:53:64
			return acc + 1
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:53:71*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:53:5
		if assertLeft2 != 2 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/ops.test.lis:53:5
			ctx.FailAssert(4, 1768, 1833, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft2)}, testkit.Operand{Label: "right", Value: lisette.Debug(2)})
		}
	}
}

func TestSmapReduceKvAndTransduce(t *testing.T) {
	defer testkit.Recover(t, 4, 1477, 1505)
	smapReduceKvAndTransduce(testkit.New(t))
}
