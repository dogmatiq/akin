package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestSingleton(t *testing.T) {
	t.Run("concrete value", func(t *testing.T) {
		set := Singleton(1)

		AssertIsMember(t, set, 1)
		AssertIsNotMember(t, set, 0)
		AssertIsNotMember(t, set, uint(1))
		AssertIsNotMember(t, set, float64(1))
	})

	for i1, c1 := range comparable {
		t.Run(c1.Name, func(t *testing.T) {
			set := Singleton(c1.Value)

			for i2, c2 := range comparable {
				t.Run(c2.Name, func(t *testing.T) {
					if i1 == i2 {
						AssertIsMember(t, set, c2.Value)
					} else {
						AssertIsNotMember(t, set, c2.Value)
					}
				})
			}
		})
	}
}
