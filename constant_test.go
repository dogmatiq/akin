package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestConstant(t *testing.T) {
	assertInvariants(t, Top)
	assertInvariants(t, Bottom)

	assertIsReduced(t, Top)
	assertIsReduced(t, Bottom)

	for _, c := range all {
		t.Run(c.Name, func(t *testing.T) {
			assertSatisfied(t, Top, c.Value)
			assertViolated(t, Bottom, c.Value)
		})
	}
}
