package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestNil(t *testing.T) {
	for _, c := range nils {
		t.Run(c.Name, func(t *testing.T) {
			AssertSatisfied(t, IsNil, c.Value)
			AssertViolated(t, IsNonNil, c.Value)
		})
	}

	for _, c := range nonNils {
		t.Run(c.Name, func(t *testing.T) {
			AssertSatisfied(t, IsNonNil, c.Value)
			AssertViolated(t, IsNil, c.Value)
		})
	}
}

func TestNil_model(t *testing.T) {
	p := To(nil)

	for _, c := range nils {
		t.Run(c.Name, func(t *testing.T) {
			AssertSatisfied(t, p, c.Value)
		})
	}

	for _, c := range nonNils {
		t.Run(c.Name, func(t *testing.T) {
			AssertViolated(t, p, c.Value)
		})
	}
}

func TestNil_uintptr(t *testing.T) {
	// The zero-valued uintptr is not technically nil, but it is "conceptually"
	// nil, so it is treated as such when using a predicate produced by a model
	// nil value, but not when using the actual [IsNil] predicate.
	AssertSatisfied(t, To(nil), uintptr(0))
	AssertViolated(t, IsNil, uintptr(0))
	AssertSatisfied(t, IsNonNil, uintptr(0))
}
