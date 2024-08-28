package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/assert"
)

func TestInhabitance(t *testing.T) {
	p := InhabitsType[int]()

	assert.Satisfied(t, p, 0)
	assert.Violated(t, p, uint(0))
	assert.Violated(t, p, float64(0))
	assert.Violated(t, p, "0")
}
