package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestTo_bool(t *testing.T) {
	t.Skip()
	type userDefined bool
	const (
		userTrue  = userDefined(true)
		userFalse = userDefined(false)
	)

	t.Run("built-in true", func(t *testing.T) {
		AssertContains(t, To(true), true)
		AssertContains(t, To(true), userTrue)
		AssertNotContains(t, To(true), false)
		AssertNotContains(t, To(true), userFalse)
	})

	t.Run("built-in false", func(t *testing.T) {
		AssertContains(t, To(false), false)
		AssertContains(t, To(false), userFalse)
		AssertNotContains(t, To(false), true)
		AssertNotContains(t, To(false), userTrue)
	})

	t.Run("user-defined true", func(t *testing.T) {
		AssertContains(t, To(userTrue), userTrue)
		AssertNotContains(t, To(userTrue), userFalse)
		AssertNotContains(t, To(userTrue), true) // user-defined spec requires types to match
		AssertNotContains(t, To(userTrue), false)
	})

	t.Run("user-defined false", func(t *testing.T) {
		AssertContains(t, To(userFalse), userFalse)
		AssertNotContains(t, To(userFalse), userTrue)
		AssertNotContains(t, To(userFalse), true)
		AssertNotContains(t, To(userFalse), false) // user-defined spec requires types to match
	})
}
