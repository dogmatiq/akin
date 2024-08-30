package assert

import (
	"testing"

	"github.com/dogmatiq/akin"
)

// Satisfied asserts that p is satisfied by v.
func Satisfied(t *testing.T, p akin.Predicate, v any) {
	t.Helper()
	assert(t, true, p, v)
}

// Violated asserts that p is not satisfied by v.
func Violated(t *testing.T, p akin.Predicate, v any) {
	t.Helper()
	assert(t, false, p, v)
}

func assert(
	t *testing.T,
	expect bool,
	p akin.Predicate,
	x any,
) {
	t.Helper()

	e := akin.Eval(p, x)

	if e.IsSatisfied == expect {
		t.Logf("\x1b[32m✔\x1b[0m %+s", e)
	} else {
		t.Errorf("\x1b[31m✘\x1b[0m %+s", e)
	}
}

// // ReducesTo asserts that p reduces to a specific predicate.
// func ReducesTo(t *testing.T, p, want akin.Predicate) {
// 	t.Helper()

// 	got := p.Reduce()

// 	if got.Is(p) {
// 		t.Fatalf("expected %s to reduce to %s but no reduction occurred", p, want)
// 	}

// 	if !got.Is(want) {
// 		t.Fatalf("expected %s to reduce to %s but got %s", p, want, got)
// 	}
// }

// IsReduced asserts that p is in its simplified form.
func IsReduced(t *testing.T, p akin.Predicate) {
	t.Helper()

	got := akin.Reduce(p)

	if !akin.Same(p, got) {
		t.Fatalf("did not expect further reduction of %s but got %s", p, got)
	}
}

// // IsOrReducesTo asserts that the p either is, or reduces to, a specific
// // predicate.
// func IsOrReducesTo(t *testing.T, got, want akin.Predicate) {
// 	t.Helper()

// 	if got.Is(want) {
// 		return
// 	}

// 	reduced := got.Reduce()
// 	if reduced.Is(want) {
// 		return
// 	}

// 	t.Fatalf("expected %s to be equivalent to (or reduce to) %s, but got %s", got, want, reduced)
// }

// // Is asserts that two predicates are equivalent.
// func Is(t *testing.T, got, want akin.Predicate) {
// 	t.Helper()

// 	if !got.Is(want) {
// 		t.Fatalf("expected %s to be equivalent to %s", got, want)
// 	}
// }

// // IsNot asserts that two predicates are not equivalent.
// func IsNot(t *testing.T, got, dontWant akin.Predicate) {
// 	t.Helper()

// 	if got.Is(dontWant) {
// 		t.Fatalf("did not expect %s to be equivalent to %s", got, dontWant)
// 	}
// }
