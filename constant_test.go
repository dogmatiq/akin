package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/assert"
	"github.com/dogmatiq/akin/internal/testcase"
)

func TestConstant(t *testing.T) {
	assert.IsReduced(t, Top)
	assert.IsReduced(t, Bottom)

	for _, c := range testcase.All {
		assert.Satisfied(t, Top, c.Value)
		assert.Violated(t, Bottom, c.Value)
	}
}
