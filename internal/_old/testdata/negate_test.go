package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/assert"
	"github.com/dogmatiq/akin/internal/testcase"
)

func TestNot(t *testing.T) {
	assert.ReducesTo(t, Not(Top), Bottom)
	assert.ReducesTo(t, Not(Bottom), Top)

	for n, x := range testcase.All {
		t.Run(n, func(t *testing.T) {
			assert.Violated(t, Not(Bottom), x)
			assert.Satisfied(t, Not(Top), x)
		})
	}
}
