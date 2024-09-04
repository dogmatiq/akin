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

	for n, x := range testcase.All {
		t.Run(n, func(t *testing.T) {
			assert.True(t, Top, x)
			assert.False(t, Bottom, x)
		})
	}
}
