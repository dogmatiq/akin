package akin_test

import "testing"

func Test_bool(t *testing.T) {
	type userDefined bool
	const (
		userTrue  = userDefined(true)
		userFalse = userDefined(false)
	)

	// t.Run("built-in true", func(t *testing.T) {
	// 	assertAkin(t, true, true)
	// 	assertAkin(t, true, userTrue)
	// 	assertNotAkin(t, true, false)
	// 	assertNotAkin(t, true, userFalse)
	// })

	// t.Run("built-in false", func(t *testing.T) {
	// 	assertAkin(t, false, false)
	// 	assertAkin(t, false, userFalse)
	// 	assertNotAkin(t, false, true)
	// 	assertNotAkin(t, false, userTrue)
	// })

	// t.Run("user-defined true", func(t *testing.T) {
	// 	assertAkin(t, userTrue, userTrue)
	// 	assertNotAkin(t, userTrue, userFalse)
	// 	assertNotAkin(t, userTrue, true) // user-defined spec requires types to match
	// 	assertNotAkin(t, userTrue, false)
	// })

	// t.Run("user-defined false", func(t *testing.T) {
	// 	assertAkin(t, userFalse, userFalse)
	// 	assertNotAkin(t, userFalse, userTrue)
	// 	assertNotAkin(t, userFalse, true)
	// 	assertNotAkin(t, userFalse, false) // user-defined spec requires types to match
	// })
}
