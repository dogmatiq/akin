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

	if e.IsSatisfied {
		t.Logf(
			"as expected, %s satisfies %s, because %s",
			reflectx.Sprint(v),
			p,
			e.Reason,
		)
	} else {
		t.Errorf(
			"expected %s to satisfy %s, but %s",
			reflectx.Sprint(v),
			p,
			e.Reason,
		)
	}

	if e.Reason == "" {
		panic("akin: no reason provided, this is a bug")
	}
}

func assertViolated(t *testing.T, p Predicate, v any) {
	t.Helper()

	e := p.Eval(v)

	if e.IsSatisfied {
		t.Errorf(
			"expected %s to violate %s, but %s",
			reflectx.Sprint(v),
			p,
			e.Reason,
		)
	} else {
		t.Logf(
			"as expected, %s violates %s because %s",
			reflectx.Sprint(v),
			p,
			e.Reason,
		)
	}

	if e.Reason == "" {
		panic("akin: no reason provided, this is a bug")
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
