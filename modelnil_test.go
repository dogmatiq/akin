package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/testcase"
)

func TestTo_nil(t *testing.T) {
	p := To(nil)

	assertInvariants(t, p)

	for _, c := range testcase.Nil {
		assertSatisfied(t, p, c.Value)
	}

	for _, c := range testcase.NonNil {
		assertViolated(t, p, c.Value)
	}

	// The zero-valued uintptr is not technically nil, but it is "conceptually"
	// nil, so it is treated as such when using a predicate produced by a model
	// nil value, but not when using the actual [IsNil] predicate.
	assertSatisfied(t, To(nil), uintptr(0))
	assertViolated(t, IsNil, uintptr(0))
	assertSatisfied(t, IsNonNil, uintptr(0))
}
