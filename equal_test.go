package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestEqualTo(t *testing.T) {
	p := Equal(1)

	assertInvariants(t, p)

	assertSatisfied(t, p, 1)

	assertViolated(t, p, 0)
	assertViolated(t, p, uint(1))
	assertViolated(t, p, float64(1))

	assertIsReduced(t, p)

	for i1, c1 := range comparable {
		t.Run(c1.Name, func(t *testing.T) {
			p := Equal(c1.Value)

			for i2, c2 := range comparable {
				t.Run(c2.Name, func(t *testing.T) {
					if i1 == i2 {
						assertSatisfied(t, p, c2.Value)
					} else {
						assertViolated(t, p, c2.Value)
					}
				})
			}
		})
	}
}
