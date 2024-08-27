package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestSingleton(t *testing.T) {
	t.Run("concrete value", func(t *testing.T) {
		p := EqualTo(1)

		AssertSatisfied(t, p, 1)
		AssertViolated(t, p, 0)
		AssertViolated(t, p, uint(1))
		AssertViolated(t, p, float64(1))
	})

	for i1, c1 := range comparable {
		t.Run(c1.Name, func(t *testing.T) {
			p := EqualTo(c1.Value)

			for i2, c2 := range comparable {
				t.Run(c2.Name, func(t *testing.T) {
					if i1 == i2 {
						AssertSatisfied(t, p, c2.Value)
					} else {
						AssertViolated(t, p, c2.Value)
					}
				})
			}
		})
	}
}
