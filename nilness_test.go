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
			assert.True(t, IsNil, x)
			assert.False(t, IsNonNil, x)
		})
	}

	for n, x := range testcase.NonNil {
		t.Run(n, func(t *testing.T) {
			assert.False(t, IsNil, x)
			assert.True(t, IsNonNil, x)
		})
	}
}
