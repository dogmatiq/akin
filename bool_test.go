package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestBool_model(t *testing.T) {
	type casual bool
	const (
		yeah = casual(true)
		nah  = casual(false)
	)

	t.Run("built-in true", func(t *testing.T) {
		p := To(true)

		assertInvariants(t, p)

		assertSatisfied(t, p, true)
		assertViolated(t, p, false)

		assertSatisfied(t, p, yeah)
		assertViolated(t, p, nah)
	})

	t.Run("built-in false", func(t *testing.T) {
		p := To(false)

		assertInvariants(t, p)

		assertSatisfied(t, p, false)
		assertViolated(t, p, true)

		assertSatisfied(t, p, nah)
		assertViolated(t, p, yeah)
	})

	t.Run("user-defined true", func(t *testing.T) {
		p := To(yeah)

		assertInvariants(t, p)

		assertSatisfied(t, p, yeah)
		assertViolated(t, p, nah)

		// model built from user-defined type requires types to match
		assertViolated(t, p, true)
		assertViolated(t, p, false)
	})

	t.Run("user-defined false", func(t *testing.T) {
		p := To(nah)

		assertInvariants(t, p)

		assertSatisfied(t, p, nah)
		assertViolated(t, p, yeah)

		// model built from user-defined type requires types to match
		assertViolated(t, p, true)
		assertViolated(t, p, false)
	})
}
