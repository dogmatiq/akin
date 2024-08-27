package akin_test

import (
	"math"
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestEquivalentTo(t *testing.T) {
	p := EquivalentTo(1)

	assertInvariants(t, p)

	assertSatisfied(t, p, 1)
	assertSatisfied(t, p, uint(1))
	assertSatisfied(t, p, float64(1))

	assertViolated(t, EquivalentTo(int8(math.MinInt8)), math.MaxInt8+1)          // "8-bit signed overflow"
	assertViolated(t, EquivalentTo(int16(math.MinInt16)), math.MaxInt16+1)       // "16-bit signed overflow"
	assertViolated(t, EquivalentTo(int32(math.MinInt32)), math.MaxInt32+1)       // "32-bit signed overflow"
	assertViolated(t, EquivalentTo(int64(math.MinInt64)), uint(math.MaxInt64+1)) // "64-bit signed overflow"
	assertViolated(t, EquivalentTo(uint8(0)), math.MaxUint8+1)                   // "8-bit unsigned overflow"
	assertViolated(t, EquivalentTo(uint16(0)), math.MaxUint16+1)                 // "16-bit unsigned overflow"
	assertViolated(t, EquivalentTo(uint32(0)), math.MaxUint32+1)                 // "32-bit unsigned overflow"
	assertViolated(t, EquivalentTo(uint64(0)), float64(math.MaxUint64+1))        // "64-bit unsigned overflow"

	for i1, c1 := range comparable {
		t.Run(c1.Name, func(t *testing.T) {
			p := EquivalentTo(c1.Value)

			for i2, c2 := range comparable {
				t.Run(c2.Name, func(t *testing.T) {
					if i1 == i2 {
						assertSatisfied(t, p, c2.Value)
					}
				})
			}
		})
	}
}
