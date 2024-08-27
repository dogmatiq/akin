package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestIsNil(t *testing.T) {
	assertInvariants(t, IsNil)
	assertInvariants(t, IsNonNil)

	assertIsReduced(t, IsNil)
	assertIsReduced(t, IsNonNil)

	for _, c := range nils {
		t.Run(c.Name, func(t *testing.T) {
			assertSatisfied(t, IsNil, c.Value)
			assertViolated(t, IsNonNil, c.Value)
		})
	}

	for _, c := range nonNils {
		t.Run(c.Name, func(t *testing.T) {
			assertSatisfied(t, IsNonNil, c.Value)
			assertViolated(t, IsNil, c.Value)
		})
	}
}
