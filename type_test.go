package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/assert"
)

func TestHasType(t *testing.T) {
	p := HasType[int]()

	assert.IsReduced(t, p)

	assert.Satisfied(t, p, 0)
	assert.Violated(t, p, uint(0))
	assert.Violated(t, p, float64(0))
	assert.Violated(t, p, "0")
}
