package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestEquality(t *testing.T) {
	p := IsEqualTo(1)

	assertInvariants(t, p)

	assertSatisfied(t, p, 1)

	assertViolated(t, p, 0)
	assertViolated(t, p, uint(1))
	assertViolated(t, p, float64(1))

	assertIsReduced(t, p)

	for i1, c1 := range comparable {
		p := IsEqualTo(c1.Value)

		for i2, c2 := range comparable {
			if i1 == i2 {
				assertSatisfied(t, p, c2.Value)
			} else {
				assertViolated(t, p, c2.Value)
			}
		}
	}
}
