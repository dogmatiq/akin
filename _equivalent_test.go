package akin_test

import (
	"math"
	"reflect"
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/assert"
	"github.com/dogmatiq/akin/internal/testcase"
)

func TestEquivalence(t *testing.T) {
	p := IsEquivalentTo(1)

	assert.Satisfied(t, p, 1)
	assert.Satisfied(t, p, uint(1))
	assert.Satisfied(t, p, float64(1))

	assert.Violated(t, IsEquivalentTo(int8(math.MinInt8)), math.MaxInt8+1)          // "8-bit signed overflow"
	assert.Violated(t, IsEquivalentTo(int16(math.MinInt16)), math.MaxInt16+1)       // "16-bit signed overflow"
	assert.Violated(t, IsEquivalentTo(int32(math.MinInt32)), math.MaxInt32+1)       // "32-bit signed overflow"
	assert.Violated(t, IsEquivalentTo(int64(math.MinInt64)), uint(math.MaxInt64+1)) // "64-bit signed overflow"
	assert.Violated(t, IsEquivalentTo(uint8(0)), math.MaxUint8+1)                   // "8-bit unsigned overflow"
	assert.Violated(t, IsEquivalentTo(uint16(0)), math.MaxUint16+1)                 // "16-bit unsigned overflow"
	assert.Violated(t, IsEquivalentTo(uint32(0)), math.MaxUint32+1)                 // "32-bit unsigned overflow"
	assert.Violated(t, IsEquivalentTo(uint64(0)), float64(math.MaxUint64+1))        // "64-bit unsigned overflow"

	for _, c := range testcase.Comparable {
		assert.Satisfied(t, IsEquivalentTo(c.Value), c.Value)
	}

	for _, numbers := range []testcase.Set{testcase.Zero, testcase.Pos, testcase.Neg} {
		_, nonComplex := numbers.Split(reflect.Value.CanComplex)

		for _, c1 := range nonComplex {
			p := IsEquivalentTo(c1.Value)

			for _, c2 := range nonComplex {
				assert.Satisfied(t, p, c2.Value)
			}
		}
	}
}
