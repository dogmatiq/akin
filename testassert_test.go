package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/reflectx"
)

func assertInvariants(t *testing.T, p Predicate) {
	t.Helper()

	if !p.Is(p) {
		t.Fatalf("%s does not compare as equal to itself", p)
	}
}

func assertSatisfied(t *testing.T, p Predicate, v any) {
	t.Helper()

	e := p.Eval(v)

	if e.Reason == "" {
		panic("akin: no reason provided, this is a bug")
	}

	if e.IsSatisfied {
		t.Logf(
			"as expected, given 𝑥 ≔ %s, the predicate %s is satisfied, because %s",
			reflectx.Sprint(v),
			p,
			e.Reason,
		)
	} else {
		t.Fatalf(
			"given 𝑥 ≔ %s, the predicate %s is unexpectedly violated, because %s",
			reflectx.Sprint(v),
			p,
			e.Reason,
		)
	}
}

func assertViolated(t *testing.T, p Predicate, v any) {
	t.Helper()

	e := p.Eval(v)

	if e.Reason == "" {
		panic("akin: no reason provided, this is a bug")
	}

	if e.IsSatisfied {
		t.Fatalf(
			"given 𝑥 ≔ %s, the predicate %s is unexpectedly satisfied, because %s",
			reflectx.Sprint(v),
			p,
			e.Reason,
		)
	} else {
		t.Logf(
			"as expected, given 𝑥 ≔ %s, the predicate %s is violated, because %s",
			reflectx.Sprint(v),
			p,
			e.Reason,
		)
	}
}

func assertReducesTo(t *testing.T, want, p Predicate) {
	t.Helper()

	got := p.Reduce()

	if got.Is(p) {
		t.Fatalf("expected %s to reduce to %s but no reduction occurred", p, want)
	}

	if !got.Is(want) {
		t.Fatalf("expected %s to reduce to %s but got %s", p, want, got)
	}
}

func assertIsReduced(t *testing.T, p Predicate) {
	t.Helper()

	got := p.Reduce()

	if !got.Is(p) {
		t.Fatalf("did not expect reduction of %s but got %s", p, got)
	}
}

func assertIsOrReducesTo(t *testing.T, want, p Predicate) {
	t.Helper()

	if p.Is(want) {
		return
	}

	got := p.Reduce()
	if got.Is(want) {
		return
	}

	t.Fatalf("expected %s to be (or reduce to) %s, but got %s", p, want, got)
}

func assertEquivalent(t *testing.T, p, q Predicate) {
	t.Helper()

	if !p.Is(q) {
		t.Fatalf("expected %s to be equal to %s", p, q)
	}
}

func assertNotEquivalent(t *testing.T, p, q Predicate) {
	t.Helper()

	if p.Is(q) {
		t.Fatalf("did not expect %s to be equal to %s", p, q)
	}
}
