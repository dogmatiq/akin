package assert

import (
	"testing"

	"github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/reflectx"
)

// Satisfied asserts that p is satisfied by v.
func Satisfied(t *testing.T, p akin.Predicate, v any) {
	t.Helper()

	e := eval(t, p, v)

	if e.IsSatisfied {
		t.Logf(
			"as expected, given 𝑥 ≔ %s, the predicate %s is satisfied, because %s",
			reflectx.Sprint(v),
			p,
			e.Description,
		)
	} else {
		t.Fatalf(
			"given 𝑥 ≔ %s, the predicate %s is unexpectedly violated, because %s",
			reflectx.Sprint(v),
			p,
			e.Description,
		)
	}
}

// Violated asserts that p is not satisfied by v.
func Violated(t *testing.T, p akin.Predicate, v any) {
	t.Helper()

	e := eval(t, p, v)

	if e.IsSatisfied {
		t.Fatalf(
			"given 𝑥 ≔ %s, the predicate %s is unexpectedly satisfied, because %s",
			reflectx.Sprint(v),
			p,
			e.Description,
		)
	} else {
		t.Logf(
			"as expected, given 𝑥 ≔ %s, the predicate %s is violated, because %s",
			reflectx.Sprint(v),
			p,
			e.Description,
		)
	}
}

// ReducesTo asserts that p reduces to a specific predicate.
func ReducesTo(t *testing.T, p, want akin.Predicate) {
	t.Helper()

	got := p.Reduce()

	if got.Is(p) {
		t.Fatalf("expected %s to reduce to %s but no reduction occurred", p, want)
	}

	if !got.Is(want) {
		t.Fatalf("expected %s to reduce to %s but got %s", p, want, got)
	}
}

// IsReduced asserts that p does not reduce, that is, it is already in its most
// reduced form.
func IsReduced(t *testing.T, p akin.Predicate) {
	t.Helper()

	got := p.Reduce()

	if !got.Is(p) {
		t.Fatalf("did not expect reduction of %s but got %s", p, got)
	}
}

// IsOrReducesTo asserts that the p either is, or reduces to, a specific
// predicate.
func IsOrReducesTo(t *testing.T, got, want akin.Predicate) {
	t.Helper()

	if got.Is(want) {
		return
	}

	reduced := got.Reduce()
	if reduced.Is(want) {
		return
	}

	t.Fatalf("expected %s to be equivalent to (or reduce to) %s, but got %s", got, want, reduced)
}

// Is asserts that two predicates are equivalent.
func Is(t *testing.T, got, want akin.Predicate) {
	t.Helper()

	if !got.Is(want) {
		t.Fatalf("expected %s to be equivalent to %s", got, want)
	}
}

// IsNot asserts that two predicates are not equivalent.
func IsNot(t *testing.T, got, dontWant akin.Predicate) {
	t.Helper()

	if got.Is(dontWant) {
		t.Fatalf("did not expect %s to be equivalent to %s", got, dontWant)
	}
}

func eval(t *testing.T, p akin.Predicate, v any) akin.Evaluation {
	t.Helper()

	if !p.Is(p) {
		t.Fatalf("%s does not compare as equivalent to itself", p)
	}

	e := p.Eval(v)

	if e.Description == "" {
		t.Fatalf("no evaluation description provided")
	}

	return e
}
