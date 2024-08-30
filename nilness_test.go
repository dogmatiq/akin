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

	for n, x := range testcase.Nil {
		t.Run(n, func(t *testing.T) {
			assert.Satisfied(t, IsNil, x)
			assert.Violated(t, IsNonNil, x)
		})
	}

	for n, x := range testcase.NonNil {
		t.Run(n, func(t *testing.T) {
			assert.Violated(t, IsNil, x)
			assert.Satisfied(t, IsNonNil, x)
		})
	}
}
