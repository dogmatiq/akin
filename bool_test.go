package akin_test

import "testing"

func Test_bool(t *testing.T) {
	type state bool
	const (
		enabled  = state(true)
		disabled = state(false)
	)

	t.Run("built-in true", func(t *testing.T) {
		assertAkin(t, true, true)
		assertAkin(t, true, enabled)
		assertNotAkin(t, true, false)
		assertNotAkin(t, true, disabled)
	})

	t.Run("built-in false", func(t *testing.T) {
		assertAkin(t, false, false)
		assertAkin(t, false, disabled)
		assertNotAkin(t, false, true)
		assertNotAkin(t, false, enabled)
	})

	t.Run("user-defined true", func(t *testing.T) {
		assertAkin(t, enabled, enabled)
		assertNotAkin(t, enabled, disabled)
		assertNotAkin(t, enabled, true) // user-defined spec requires types to match
		assertNotAkin(t, enabled, false)
	})

	t.Run("user-defined false", func(t *testing.T) {
		assertAkin(t, disabled, disabled)
		assertNotAkin(t, disabled, enabled)
		assertNotAkin(t, disabled, true)
		assertNotAkin(t, disabled, false) // user-defined spec requires types to match
	})
}
