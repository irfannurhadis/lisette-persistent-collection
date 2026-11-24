package collections

import (
	"slices"
	"strings"
	"testing"

	lisette "github.com/ivov/lisette/prelude"
	"github.com/ivov/lisette/prelude/testkit"
)

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:7:4
func eq[A comparable](a, b []A) bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:8:5
	return slices.Equal(a, b)
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:12:4
func identityLaw(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:13:5
	xs := []int{1, 2, 3}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:14:5
	if !eq(Run(Identity[int](), xs), xs) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:14:5
		ctx.FailAssert(9, 367, 394, "bare", "assertion failed")
	}
}

func TestIdentityLaw(t *testing.T) {
	defer testkit.Recover(t, 9, 316, 328)
	identityLaw(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:18:4
func mapLaw(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:19:5
	xs := []int{1, 2, 3}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:20:5
	if !eq(Run(Map(func(x int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:20:32
		return x + 1
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:20:37*/), xs), []int{2, 3, 4}) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:20:5
		ctx.FailAssert(9, 455, 498, "bare", "assertion failed")
	}
}

func TestMapLaw(t *testing.T) {
	defer testkit.Recover(t, 9, 409, 416)
	mapLaw(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:24:4
func filterLaw(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:25:5
	xs := []int{1, 2, 3, 4}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:26:5
	if !eq(Run(Filter(func(x int) bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:26:35
		return x > 2
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:26:40*/), xs), []int{3, 4}) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:26:5
		ctx.FailAssert(9, 565, 608, "bare", "assertion failed")
	}
}

func TestFilterLaw(t *testing.T) {
	defer testkit.Recover(t, 9, 513, 523)
	filterLaw(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:30:4
func takeDropLaw(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:31:5
	xs := []int{1, 2, 3, 4}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:32:5
	if !eq(Run(Take[int](2), xs), []int{1, 2}) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:32:5
		ctx.FailAssert(9, 678, 706, "bare", "assertion failed")
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:33:5
	if !eq(Run(Drop[int](2), xs), []int{3, 4}) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:33:5
		ctx.FailAssert(9, 718, 746, "bare", "assertion failed")
	}
}

func TestTakeDropLaw(t *testing.T) {
	defer testkit.Recover(t, 9, 623, 636)
	takeDropLaw(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:37:4
func compositionIsAssociative(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:38:5
	xs := []int{1, 2, 3, 4}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:39:5
	f := Map(func(x int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:39:26
		return x + 1
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:39:31*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:40:5
	g := Filter(func(x int) bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:40:29
		return x > 2
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:40:34*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:41:5
	h := Take[int](1)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:42:5
	left := Run(Compose(Compose(f, g), h), xs)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:43:5
	right := Run(Compose(f, Compose(g, h)), xs)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:44:5
	if !eq(left, right) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:44:5
		ctx.FailAssert(9, 1017, 1032, "bare", "assertion failed")
	}
}

func TestCompositionIsAssociative(t *testing.T) {
	defer testkit.Recover(t, 9, 761, 787)
	compositionIsAssociative(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:48:4
func fusionEqualsMaterialized(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:50:5
	xs := []int{1, 2, 3, 4, 5}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:51:5
	fused := Run(Compose(Map(func(x int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:51:42
		return x * 2
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:51:47*/), Filter(func(x int) bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:51:66
		return x > 4
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:51:71*/)), xs)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:52:5
	materialized := FilterCollect(MapCollect(xs, func(x int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:52:64
		return x * 2
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:52:69*/), func(x int) bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:52:81
		return x > 4
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:52:86*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:53:5
	if !eq(fused, materialized) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:53:5
		ctx.FailAssert(9, 1361, 1384, "bare", "assertion failed")
	}
}

func TestFusionEqualsMaterialized(t *testing.T) {
	defer testkit.Recover(t, 9, 1047, 1073)
	fusionEqualsMaterialized(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:57:4
func foldIsOneAccumulator(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:58:5
	xs := []int{1, 2, 3}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:59:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:59:5
		assertLeft := Fold(Identity[int](), xs, 0, func(acc, a int) int {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:59:55
			return acc + a
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:59:62*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:59:5
		if assertLeft != 6 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:59:5
			ctx.FailAssert(9, 1461, 1517, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft)}, testkit.Operand{Label: "right", Value: lisette.Debug(6)})
		}
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:60:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:60:5
		assertLeft2 := SumCollect(xs)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:60:5
		if assertLeft2 != 6 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:60:5
			ctx.FailAssert(9, 1529, 1549, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft2)}, testkit.Operand{Label: "right", Value: lisette.Debug(6)})
		}
	}
}

func TestFoldIsOneAccumulator(t *testing.T) {
	defer testkit.Recover(t, 9, 1399, 1422)
	foldIsOneAccumulator(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:64:4
func findShortCircuits(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:65:5
	xs := []int{1, 2, 3, 4}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:66:5
	found := lisette.OptionFromCommaOk[int](Find(Identity[int](), xs, func(x int) bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:66:47
		return x > 2
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:66:52*/))
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:67:5
	if found.Tag != lisette.OptionSome {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:67:32
		return
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:67:5
	v := found.SomeVal
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:68:5
	if v != 3 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:68:5
		ctx.FailAssert(9, 1718, 1724, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(v)}, testkit.Operand{Label: "right", Value: lisette.Debug(3)})
	}
}

func TestFindShortCircuits(t *testing.T) {
	defer testkit.Recover(t, 9, 1564, 1583)
	findShortCircuits(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:72:4
func collectionWrappers(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:73:5
	xs := []int{1, 2, 3, 4}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:74:5
	if !eq(FilterCollect(xs, func(x int) bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:74:43
		return x%2 == 0
	} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:74:53*/), []int{2, 4}) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:74:5
		ctx.FailAssert(9, 1800, 1851, "bare", "assertion failed")
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:75:5
	if !eq(TakeCollect(xs, 2), []int{1, 2}) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:75:5
		ctx.FailAssert(9, 1863, 1894, "bare", "assertion failed")
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:76:5
	if !eq(DropCollect(xs, 3), []int{4}) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:76:5
		ctx.FailAssert(9, 1906, 1934, "bare", "assertion failed")
	}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:77:5
	{
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:77:5
		assertLeft := CountMatching(xs, func(x int) bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:77:40
			return x > 2
		} /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:77:45*/)
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:77:5
		if assertLeft != 2 {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:77:5
			ctx.FailAssert(9, 1946, 1985, "relation", "expected ==", testkit.Operand{Label: "left", Value: lisette.Debug(assertLeft)}, testkit.Operand{Label: "right", Value: lisette.Debug(2)})
		}
	}
}

func TestCollectionWrappers(t *testing.T) {
	defer testkit.Recover(t, 9, 1739, 1758)
	collectionWrappers(testkit.New(t))
}

//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:81:4
func selectIdsPipeline(ctx testkit.TestContext) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:82:5
	ids := []string{"a1", "b2", "c3"}
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:83:5
	if !eq(SelectIds(ids, func(id string) bool {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:83:44
		return strings.Contains(id, "b")
	}, /*line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:83:60*/ 2), []string{"b2"}) {
//line /home/irfan/consultancy/lisette-persistent-collection/src/collections/transducer.test.lis:83:5
		ctx.FailAssert(9, 2068, 2129, "bare", "assertion failed")
	}
}

func TestSelectIdsPipeline(t *testing.T) {
	defer testkit.Recover(t, 9, 2000, 2019)
	selectIdsPipeline(testkit.New(t))
}
