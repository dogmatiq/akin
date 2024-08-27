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
		AssertIsMember(t, To(true), true)
		AssertIsMember(t, To(true), userTrue)
		AssertIsNotMember(t, To(true), false)
		AssertIsNotMember(t, To(true), userFalse)
	})

	t.Run("built-in false", func(t *testing.T) {
		AssertIsMember(t, To(false), false)
		AssertIsMember(t, To(false), userFalse)
		AssertIsNotMember(t, To(false), true)
		AssertIsNotMember(t, To(false), userTrue)
	})

	t.Run("user-defined true", func(t *testing.T) {
		AssertIsMember(t, To(userTrue), userTrue)
		AssertIsNotMember(t, To(userTrue), userFalse)
		AssertIsNotMember(t, To(userTrue), true) // user-defined spec requires types to match
		AssertIsNotMember(t, To(userTrue), false)
	})

	t.Run("user-defined false", func(t *testing.T) {
		AssertIsMember(t, To(userFalse), userFalse)
		AssertIsNotMember(t, To(userFalse), userTrue)
		AssertIsNotMember(t, To(userFalse), true)
		AssertIsNotMember(t, To(userFalse), false) // user-defined spec requires types to match
	})
}
