package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestBool_model(t *testing.T) {
	type userDefined bool
	const (
		userTrue  = userDefined(true)
		userFalse = userDefined(false)
	)

	t.Run("built-in true", func(t *testing.T) {
		AssertSatisfied(t, To(true), true)
		AssertSatisfied(t, To(true), userTrue)
		AssertViolated(t, To(true), false)
		AssertViolated(t, To(true), userFalse)
	})

	t.Run("built-in false", func(t *testing.T) {
		AssertSatisfied(t, To(false), false)
		AssertSatisfied(t, To(false), userFalse)
		AssertViolated(t, To(false), true)
		AssertViolated(t, To(false), userTrue)
	})

	t.Run("user-defined true", func(t *testing.T) {
		AssertSatisfied(t, To(userTrue), userTrue)
		AssertViolated(t, To(userTrue), userFalse)
		AssertViolated(t, To(userTrue), true) // user-defined spec requires types to match
		AssertViolated(t, To(userTrue), false)
	})

	t.Run("user-defined false", func(t *testing.T) {
		AssertSatisfied(t, To(userFalse), userFalse)
		AssertViolated(t, To(userFalse), userTrue)
		AssertViolated(t, To(userFalse), true)
		AssertViolated(t, To(userFalse), false) // user-defined spec requires types to match
	})
}
