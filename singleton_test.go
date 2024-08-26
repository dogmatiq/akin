package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestSingleton(t *testing.T) {
	t.Run("concrete value", func(t *testing.T) {
		set := Singleton(1)

		AssertContains(t, set, 1)
		AssertNotContains(t, set, 0)
		AssertNotContains(t, set, uint(1))
		AssertNotContains(t, set, float64(1))
	})

	t.Run("nil interface", func(t *testing.T) {
		set := Singleton(nil)

		AssertContains(t, set, nil)
		AssertNotContains(t, set, (*int)(nil))
	})
}
