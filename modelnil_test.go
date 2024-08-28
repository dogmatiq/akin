package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/assert"
	"github.com/dogmatiq/akin/internal/testcase"
)

func TestTo_nil(t *testing.T) {
	p := To(nil)

	for _, c := range testcase.Nil {
		assert.Satisfied(t, p, c.Value)
	}

	for _, c := range testcase.NonNil {
		assert.Violated(t, p, c.Value)
	}

	// The zero-valued uintptr is not technically nil, but it is "conceptually"
	// nil, so it is treated as such when using a predicate produced by a model
	// nil value, but not when using the actual [IsNil] predicate.
	assert.Satisfied(t, To(nil), uintptr(0))
	assert.Violated(t, IsNil, uintptr(0))
	assert.Satisfied(t, IsNonNil, uintptr(0))
}
