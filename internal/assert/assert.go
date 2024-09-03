package assert

import (
	"testing"

	"github.com/dogmatiq/akin"
)

// True asserts that 𝑷❨𝒙❩ evaluates to true.
func True(t *testing.T, p akin.Predicate, v any) {
	t.Helper()
	assert(t, akin.True, p, v)
}

// False asserts that 𝑷❨𝒙❩ evaluates to false.
func False(t *testing.T, p akin.Predicate, v any) {
	t.Helper()
	assert(t, akin.False, p, v)
}

func assert(
	t *testing.T,
	expect akin.Truth,
	p akin.Predicate,
	x any,
) {
	t.Helper()

	result, rationale := akin.Eval(p, x)

	if result == expect {
		t.Logf("\x1b[32m✔\x1b[0m %+s", rationale)
	} else {
		t.Errorf("\x1b[31m✘\x1b[0m %+s", rationale)
	}
}

// // ReducesTo asserts that p reduces to a specific predicate.
// func ReducesTo(t *testing.T, p, want akin.Predicate) {
// 	t.Helper()

// 	got := akin.Reduce(p)

// 	if akin.Same(got, p) {
// 		t.Fatalf("expected %s to reduce to %s but no reduction occurred", p, want)
// 	}

// 	if !akin.Same(got, want) {
// 		t.Fatalf("expected %s to reduce to %s but got %s", p, want, got)
// 	}
// }

// // IsReduced asserts that p is in its simplified form.
// func IsReduced(t *testing.T, p akin.Predicate) {
// 	t.Helper()

// 	got := akin.Reduce(p)

// 	if !akin.Same(p, got) {
// 		t.Fatalf("did not expect further reduction of %s but got %s", p, got)
// 	}
// }

// // // IsOrReducesTo asserts that the p either is, or reduces to, a specific
// // // predicate.
// // func IsOrReducesTo(t *testing.T, got, want akin.Predicate) {
// // 	t.Helper()

// // 	if got.Is(want) {
// // 		return
// // 	}

// // 	reduced := got.Reduce()
// // 	if reduced.Is(want) {
// // 		return
// // 	}

// // 	t.Fatalf("expected %s to be equivalent to (or reduce to) %s, but got %s", got, want, reduced)
// // }

// // // Is asserts that two predicates are equivalent.
// // func Is(t *testing.T, got, want akin.Predicate) {
// // 	t.Helper()

// // 	if !got.Is(want) {
// // 		t.Fatalf("expected %s to be equivalent to %s", got, want)
// // 	}
// // }

// // // IsNot asserts that two predicates are not equivalent.
// // func IsNot(t *testing.T, got, dontWant akin.Predicate) {
// // 	t.Helper()

// // 	if got.Is(dontWant) {
// // 		t.Fatalf("did not expect %s to be equivalent to %s", got, dontWant)
// // 	}
// // }
