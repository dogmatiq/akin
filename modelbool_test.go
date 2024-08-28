package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/assert"
)

func TestTo_bool(t *testing.T) {
	type casual bool
	const (
		yeah = casual(true)
		nah  = casual(false)
	)

	t.Run("built-in true", func(t *testing.T) {
		p := To(true)

		assert.Satisfied(t, p, true)
		assert.Violated(t, p, false)

		assert.Satisfied(t, p, yeah)
		assert.Violated(t, p, nah)
	})

	t.Run("built-in false", func(t *testing.T) {
		p := To(false)

		assert.Satisfied(t, p, false)
		assert.Violated(t, p, true)

		assert.Satisfied(t, p, nah)
		assert.Violated(t, p, yeah)
	})

	t.Run("user-defined true", func(t *testing.T) {
		p := To(yeah)

		assert.Satisfied(t, p, yeah)
		assert.Violated(t, p, nah)

		// model built from user-defined type requires types to match
		assert.Violated(t, p, true)
		assert.Violated(t, p, false)
	})

	t.Run("user-defined false", func(t *testing.T) {
		p := To(nah)

		assert.Satisfied(t, p, nah)
		assert.Violated(t, p, yeah)

		// model built from user-defined type requires types to match
		assert.Violated(t, p, true)
		assert.Violated(t, p, false)
	})
}
