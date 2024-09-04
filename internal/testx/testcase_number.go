package testx

var (
	// NumericCases is the set of all cases involving numeric values.
	NumericCases = Union(
		ZeroCases,
		PositiveCases,
		NegativeCases,
	)

	// ZeroCases is the set of test cases for numbers with a zero value.
	ZeroCases = Cases{
		{"zero int", int(0)},
		{"zero int8", int8(0)},
		{"zero int16", int16(0)},
		{"zero int32", int32(0)},
		{"zero int64", int64(0)},
		{"zero uint", uint(0)},
		{"zero uint8", uint8(0)},
		{"zero uint16", uint16(0)},
		{"zero uint32", uint32(0)},
		{"zero uint64", uint64(0)},
		{"zero float32", float32(0)},
		{"zero float64", float64(0)},
		{"zero complex64", complex64(0)},
		{"zero complex128", complex128(0)},
		{"zero uintptr", uintptr(0)},
	}

	// PositiveCases is the set of cases for numbers with a positive value.
	//
	// All cases have the same numeric value of 1, regardless of type.
	PositiveCases = Cases{
		{"positive int", int(1)},
		{"positive int8", int8(1)},
		{"positive int16", int16(1)},
		{"positive int32", int32(1)},
		{"positive int64", int64(1)},
		{"positive uint", uint(1)},
		{"positive uint8", uint8(1)},
		{"positive uint16", uint16(1)},
		{"positive uint32", uint32(1)},
		{"positive uint64", uint64(1)},
		{"positive float32", float32(1)},
		{"positive float64", float64(1)},
		{"positive complex64", complex(float32(1), 0)},
		{"positive complex128", complex(float64(1), 0)},
		{"positive uintptr", uintptr(1)},
	}

	// NegativeCases is the set of cases for numbers with a negative value.
	//
	// All cases have the same numeric value of -1, regardless of type.
	NegativeCases = Cases{
		{"negative int", int(-1)},
		{"negative int8", int8(-1)},
		{"negative int16", int16(-1)},
		{"negative int32", int32(-1)},
		{"negative int64", int64(-1)},
		{"negative float32", float32(-1)},
		{"negative float64", float64(-1)},
		{"negative complex64", complex(float32(-1), 0)},
		{"negative complex128", complex(float64(-1), 0)},
	}
)
