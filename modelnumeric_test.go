package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	"github.com/dogmatiq/akin/internal/testcase"
)

func TestTo_numeric(t *testing.T) {
	type (
		userInt        int
		userInt8       int8
		userInt16      int16
		userInt32      int32
		userInt64      int64
		userUint       uint
		userUint8      uint8
		userUint16     uint16
		userUint32     uint32
		userUint64     uint64
		userFloat32    float32
		userFloat64    float64
		userComplex64  complex64
		userComplex128 complex128
		userUintptr    uintptr
	)

	for _, v := range []any{
		userInt(1), userInt8(1), userInt16(1), userInt32(1), userInt64(1),
		userUint(1), userUint8(1), userUint16(1), userUint32(1), userUint64(1),
		userFloat32(1), userFloat64(1),
		userComplex64(1), userComplex128(1),
		userUintptr(1),
	} {
		assertIsOrReducesTo(t, IsEqualTo(v), To(v))
	}

	for _, cases := range []testcase.Set{testcase.Zero, testcase.Pos, testcase.Neg} {
		for _, c1 := range cases {
			p := To(c1.Value)
			for _, c2 := range cases {
				assertSatisfied(t, p, c2.Value)
			}
		}
	}
}
