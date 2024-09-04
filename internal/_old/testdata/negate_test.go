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

	for _, c := range testcase.All {
		t.Run(c.Name, func(t *testing.T) {
			assert.Violated(t, Not(Bottom), c.X)
			assert.Satisfied(t, Not(Top), c.X)
		})
	}
}
