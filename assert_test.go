package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func assertSimplifiesTo(t *testing.T, want, p Predicate) {
	t.Helper()

	got, ok := p.Simplify()

	if !got.Is(want) || !ok {
		t.Fatalf("expected %s to be simplified to %s, but Simplify() returned (%s, %t)", p, want, got, ok)
	}
}

func assertIsSimple(t *testing.T, p Predicate) {
	t.Helper()

	got, ok := p.Simplify()

	if !got.Is(p) || ok {
		t.Fatalf("expected %s to be the simplest representation, but Simplify() returned (%s, %t)", p, got, ok)
	}
}
