package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/testcase"
)

func TestNilness(t *testing.T) {
	assertInvariants(t, IsNil)
	assertInvariants(t, IsNonNil)

	assertIsReduced(t, IsNil)
	assertIsReduced(t, IsNonNil)

	for _, c := range testcase.Nil {
		assertSatisfied(t, IsNil, c.Value)
		assertViolated(t, IsNonNil, c.Value)
	}

	for _, c := range testcase.NonNil {
		assertSatisfied(t, IsNonNil, c.Value)
		assertViolated(t, IsNil, c.Value)
	}
}
