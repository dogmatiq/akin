package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestAnyOf(t *testing.T) {
	for _, c := range all {
		t.Run(c.Name, func(t *testing.T) {
			AssertSatisfied(t, AnyOf(Anything, Nothing), c.Value)
			AssertSatisfied(t, AnyOf(Anything, EqualTo(1)), c.Value)
			AssertViolated(t, AnyOf(Nothing, Nothing), c.Value)
		})
	}

	p := AnyOf(
		EqualTo(1),
		EqualTo(2),
	)

	AssertSatisfied(t, p, 1)
	AssertSatisfied(t, p, 2)
	AssertViolated(t, p, 3)
}
