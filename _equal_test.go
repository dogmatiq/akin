package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/assert"
	"github.com/dogmatiq/akin/internal/testcase"
)

func TestEquality(t *testing.T) {
	p := IsEqualTo(1)

	assert.Satisfied(t, p, 1)

	assert.Violated(t, p, 0)
	assert.Violated(t, p, uint(1))
	assert.Violated(t, p, float64(1))

	assert.IsReduced(t, p)

	for i1, c1 := range testcase.Comparable {
		p := IsEqualTo(c1.Value)

		for i2, c2 := range testcase.Comparable {
			if i1 == i2 {
				assert.Satisfied(t, p, c2.Value)
			} else {
				assert.Violated(t, p, c2.Value)
			}
		}
	}
}
