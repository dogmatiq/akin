package akin_test

import (
	"math"
	"testing"

	. "github.com/dogmatiq/akin"
)

// import (
// 	"math"
// 	"testing"
// )

func TestTo_numeric(t *testing.T) {
	t.Run("user-defined types require an exact type match", func(t *testing.T) {
		type userDefined int
		AssertContains(t, To(1), userDefined(1))
		AssertNotContains(t, To(userDefined(1)), 1)
	})

	t.Run("lossless conversions", func(t *testing.T) {
		cases := []struct {
			Name   string
			Values []any
		}{
			{
				Name: "zero numeric values",
				Values: []any{
					int(0), int8(0), int16(0), int32(0), int64(0),
					uint(0), uint8(0), uint16(0), uint32(0), uint64(0),
					float32(0), float64(0),
				},
			},
			{
				Name: "positive numeric values",
				Values: []any{
					int(1), int8(1), int16(1), int32(1), int64(1),
					uint(1), uint8(1), uint16(1), uint32(1), uint64(1),
					float32(1), float64(1),
				},
			},
			{
				Name: "negative numeric values",
				Values: []any{
					int(-1), int8(-1), int16(-1), int32(-1), int64(-1),
					float32(-1), float64(-1),
				},
			},
			{
				Name: "positive infinity",
				Values: []any{
					float32(math.Inf(1)), float64(math.Inf(1)),
				},
			},
			{
				Name: "negative infinity",
				Values: []any{
					float32(math.Inf(-1)), float64(math.Inf(-1)),
				},
			},
		}

		for _, c := range cases {
			t.Run(c.Name, func(t *testing.T) {
				for _, a := range c.Values {
					set := To(a)
					for _, b := range c.Values {
						AssertContains(t, set, b)
					}
				}
			})
		}
	})

	t.Run("lossy conversions", func(t *testing.T) {
		AssertNotContains(t, To(int8(math.MinInt8)), math.MaxInt8+1)          // "8-bit signed overflow"
		AssertNotContains(t, To(int16(math.MinInt16)), math.MaxInt16+1)       // "16-bit signed overflow"
		AssertNotContains(t, To(int32(math.MinInt32)), math.MaxInt32+1)       // "32-bit signed overflow"
		AssertNotContains(t, To(int64(math.MinInt64)), uint(math.MaxInt64+1)) // "64-bit signed overflow"
		AssertNotContains(t, To(uint8(0)), math.MaxUint8+1)                   // "8-bit unsigned overflow"
		AssertNotContains(t, To(uint16(0)), math.MaxUint16+1)                 // "16-bit unsigned overflow"
		AssertNotContains(t, To(uint32(0)), math.MaxUint32+1)                 // "32-bit unsigned overflow"
		AssertNotContains(t, To(uint64(0)), float64(math.MaxUint64+1))        // "64-bit unsigned overflow"
	})
}
