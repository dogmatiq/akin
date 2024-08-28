package reflectx

import (
	"reflect"
)

// IsNeg returns true if v is a negative number.
func IsNeg(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() < 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() < 0
	case reflect.Float32, reflect.Float64:
		return v.Float() < 0
	default:
		return false
	}
}

// ToFloat64 returns v as a float64, if the conversion can be performed without
// losing information.
func ToFloat64(v reflect.Value) (float64, bool) {
	if v.CanFloat() {
		return v.Float(), true
	}

	if v.CanInt() {
		n := v.Int()
		f := float64(n)
		return f, int64(f) == n
	}

	if v.CanUint() {
		n := v.Uint()
		f := float64(n)
		return f, uint64(f) == n
	}

	if v.CanComplex() {
		c := v.Complex()
		i := imag(c)
		r := real(c)
		return r, i == 0
	}

	return 0, false
}

// ToComplex128 returns v as a complex128, if the conversion can be performed
// without losing information.
func ToComplex128(v reflect.Value) (complex128, bool) {
	if v.CanComplex() {
		return v.Complex(), true
	}

	if r, ok := ToFloat64(v); ok {
		return complex(r, 0), true
	}

	return 0, false
}
