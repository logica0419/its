package its_test

import (
	"github.com/youta-t/its"
)

func ExampleSlice() {
	// pass
	its.Slice(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2, 3},
	).OrError(t)

	// fail. Order is matter.
	its.Slice(
		its.EqEq(2), its.EqEq(1), its.EqEq(3),
	).Match(
		[]int{1, 2, 3},
	).OrError(t)

	// fail. Actual has not enough 3.
	its.Slice(
		its.EqEq(1), its.EqEq(2), its.EqEq(3), its.EqEq(3),
	).Match(
		[]int{1, 2, 3},
	).OrError(t)

	// fail. Actual has too much 3.
	its.Slice(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2, 3, 3},
	).OrError(t)

	// Output:
	// ✘ []int{ ... (len: /* got */ 3, /* want */ 3; +1, -1)		--- @ ./slice_test.go:16
	//     ✘ - /* got */ ?? == /* want */ 2		--- @ ./slice_test.go:17
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:17
	//     ✘ + /* got */ 2
	//     ✔ /* got */ 3 == /* want */ 3		--- @ ./slice_test.go:17
	//
	// ✘ []int{ ... (len: /* got */ 3, /* want */ 4; +0, -1)		--- @ ./slice_test.go:23
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:24
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_test.go:24
	//     ✔ /* got */ 3 == /* want */ 3		--- @ ./slice_test.go:24
	//     ✘ - /* got */ ?? == /* want */ 3		--- @ ./slice_test.go:24
	//
	// ✘ []int{ ... (len: /* got */ 4, /* want */ 3; +1, -0)		--- @ ./slice_test.go:30
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:31
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_test.go:31
	//     ✔ /* got */ 3 == /* want */ 3		--- @ ./slice_test.go:31
	//     ✘ + /* got */ 3
}

func ExampleSliceUnordered() {
	// pass
	its.SliceUnordered(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2, 3},
	).OrError(t)

	// pass. order is not matter.
	its.SliceUnordered(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{3, 1, 2},
	).OrError(t)

	// fail. there is an extra item 42.
	its.SliceUnordered(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2, 3, 42},
	).OrError(t)

	// fail. 3 is missing.
	its.SliceUnordered(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2},
	).OrError(t)

	// Output:
	// ✘ []int{ ... (unordered; len: /* want */ 3, /* got */ 4; +1, -0)		--- @ ./slice_test.go:72
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:73
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_test.go:73
	//     ✔ /* got */ 3 == /* want */ 3		--- @ ./slice_test.go:73
	//     ✘ + /* got */ 42
	//
	// ✘ []int{ ... (unordered; len: /* want */ 3, /* got */ 2; +0, -1)		--- @ ./slice_test.go:79
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:80
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_test.go:80
	//     ✘ - /* got */ ?? == /* want */ 3		--- @ ./slice_test.go:80
}

func ExampleSliceUnorderedContaining() {
	// pass
	its.SliceUnorderedContaining(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2, 3},
	).OrError(t)

	// pass. order is not matter.
	its.SliceUnorderedContaining(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{3, 1, 2},
	).OrError(t)

	// pass. extra item is okay.
	its.SliceUnorderedContaining(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2, 3, 42},
	).OrError(t)

	// fail. 3 is missing.
	its.SliceUnorderedContaining(
		its.EqEq(1), its.EqEq(2), its.EqEq(3),
	).Match(
		[]int{1, 2},
	).OrError(t)

	// Output:
	// ✘ []int{ ... (unordered, contain; len: /* got */ 2, /* want */ 3; -1)		--- @ ./slice_test.go:121
	//     ✔ /* got */ 1 == /* want */ 1		--- @ ./slice_test.go:122
	//     ✔ /* got */ 2 == /* want */ 2		--- @ ./slice_test.go:122
	//     ✘ - /* got */ ?? == /* want */ 3		--- @ ./slice_test.go:122
}
