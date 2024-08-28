package akin_test

import (
	"math"
	"reflect"
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestEquivalence(t *testing.T) {
	p := IsEquivalentTo(1)

	assertInvariants(t, p)

	assertSatisfied(t, p, 1)
	assertSatisfied(t, p, uint(1))
	assertSatisfied(t, p, float64(1))

	assertViolated(t, IsEquivalentTo(int8(math.MinInt8)), math.MaxInt8+1)          // "8-bit signed overflow"
	assertViolated(t, IsEquivalentTo(int16(math.MinInt16)), math.MaxInt16+1)       // "16-bit signed overflow"
	assertViolated(t, IsEquivalentTo(int32(math.MinInt32)), math.MaxInt32+1)       // "32-bit signed overflow"
	assertViolated(t, IsEquivalentTo(int64(math.MinInt64)), uint(math.MaxInt64+1)) // "64-bit signed overflow"
	assertViolated(t, IsEquivalentTo(uint8(0)), math.MaxUint8+1)                   // "8-bit unsigned overflow"
	assertViolated(t, IsEquivalentTo(uint16(0)), math.MaxUint16+1)                 // "16-bit unsigned overflow"
	assertViolated(t, IsEquivalentTo(uint32(0)), math.MaxUint32+1)                 // "32-bit unsigned overflow"
	assertViolated(t, IsEquivalentTo(uint64(0)), float64(math.MaxUint64+1))        // "64-bit unsigned overflow"

	for _, c := range comparable {
		assertSatisfied(t, IsEquivalentTo(c.Value), c.Value)
	}

	for _, numbers := range []testCases{zeroNumbers, positiveNumbers, negativeNumbers} {
		_, nonComplex := numbers.Split(reflect.Value.CanComplex)

		for _, c1 := range nonComplex {
			p := IsEquivalentTo(c1.Value)

			for _, c2 := range nonComplex {
				assertSatisfied(t, p, c2.Value)
			}
		}
	}
}
