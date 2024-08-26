package akin_test

import (
	"slices"
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestExtremum(t *testing.T) {
	values := append(
		slices.Clone(nils),
		true, false,
		int(0), int8(0), int16(0), int32(0), int64(0),
		uint(0), uint8(0), uint16(0), uint32(0), uint64(0),
		float32(0), float64(0),
		complex64(0), complex128(0),
		uintptr(0),
		"",
		struct{}{},
	)

	for _, v := range values {
		AssertContains(t, Universe, v)
		AssertNotContains(t, Empty, v)
	}
}
