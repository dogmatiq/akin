package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/assert"
)

func TestOr(t *testing.T) {
	p1 := IsEqualTo(1)
	p2 := IsEqualTo(2)

	unary := Or(p1)
	binary := Or(p1, p2)
	redundant := Or(p1, p1)

	assert.Satisfied(t, binary, 1)
	assert.Satisfied(t, binary, 2)
	assert.Violated(t, binary, 3)

	assert.Is(t, Or(p2, p1), binary)
	assert.IsNot(t, binary, unary)

	assert.ReducesTo(t, Or(), Bottom)
	assert.ReducesTo(t, Or(Bottom), Bottom)

	assert.ReducesTo(t, Or(Top), Top)
	assert.ReducesTo(t, Or(Top, p1), Top)
	assert.ReducesTo(t, Or(p1, Top), Top)

	assert.ReducesTo(t, unary, p1)
	assert.ReducesTo(t, redundant, p1)
	assert.IsReduced(t, binary)

	assert.ReducesTo(t, Or(p1, redundant), p1)
	assert.ReducesTo(t, Or(p2, unary), binary)
	assert.ReducesTo(t, Or(p1, binary), binary)
}
