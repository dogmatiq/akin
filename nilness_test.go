package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestNilness(t *testing.T) {
	assertInvariants(t, IsNil)
	assertInvariants(t, IsNonNil)

	assertIsReduced(t, IsNil)
	assertIsReduced(t, IsNonNil)

	for _, c := range nils {
		assertSatisfied(t, IsNil, c.Value)
		assertViolated(t, IsNonNil, c.Value)
	}

	for _, c := range nonNils {
		assertSatisfied(t, IsNonNil, c.Value)
		assertViolated(t, IsNil, c.Value)
	}
}
