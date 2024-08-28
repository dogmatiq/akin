package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/assert"
	"github.com/dogmatiq/akin/internal/testcase"
)

func TestNilness(t *testing.T) {
	assert.IsReduced(t, IsNil)
	assert.IsReduced(t, IsNonNil)

	for _, c := range testcase.Nil {
		assert.Satisfied(t, IsNil, c.Value)
		assert.Violated(t, IsNonNil, c.Value)
	}

	for _, c := range testcase.NonNil {
		assert.Satisfied(t, IsNonNil, c.Value)
		assert.Violated(t, IsNil, c.Value)
	}
}
