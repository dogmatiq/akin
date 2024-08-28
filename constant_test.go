package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/testcase"
)

func TestConstant(t *testing.T) {
	assertInvariants(t, Top)
	assertInvariants(t, Bottom)

	assertIsReduced(t, Top)
	assertIsReduced(t, Bottom)

	for _, c := range testcase.All {
		assertSatisfied(t, Top, c.Value)
		assertViolated(t, Bottom, c.Value)
	}
}
