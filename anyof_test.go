package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestAnyOf(t *testing.T) {
	for _, c := range comparable {
		t.Run(c.Name, func(t *testing.T) {
			AssertSatisfied(t, AnyOf(EqualTo(c.Value)), c.Value)
			AssertSatisfied(t, AnyOf(EqualTo(c.Value), EqualTo(1)), c.Value)

			AssertSatisfied(t, AnyOf(Anything, EqualTo(1)), c.Value)
			AssertSatisfied(t, AnyOf(Anything, Nothing), c.Value)

			AssertViolated(t, AnyOf(Nothing, Nothing), c.Value)
			AssertViolated(t, AnyOf(Nothing), c.Value)
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

func TestAnyOf_Simplify(t *testing.T) {
	assertSimplifiesTo(t, Anything, AnyOf(Anything, EqualTo(1)))
	assertSimplifiesTo(t, Nothing, AnyOf())
	assertSimplifiesTo(t, Nothing, AnyOf(Nothing))
}
