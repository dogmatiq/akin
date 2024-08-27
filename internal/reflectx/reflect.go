package reflectx

import (
	"fmt"
	"reflect"
)

// ValueOf returns the [reflect.Value] of v.
//
// It is similar to [reflect.ValueOf], but it does not return a
// [reflect.Invalid] value when passed nil. Instead, it returns a
// [reflect.Value] representing a nil [any] interface.
func ValueOf(v any) reflect.Value {
	if r := reflect.ValueOf(v); r.IsValid() {
		return r
	}
	if r := reflect.ValueOf(&v).Elem(); r.IsValid() {
		return r
	}
	panic(fmt.Sprintf("cannot reflect on value of %v", v))
}

// IsNil returns true if v is nil.
//
// It is similar to [reflect.Value.IsNil], but it works with any kind of value.
func IsNil(v reflect.Value) bool {
	switch v.Kind() {
	default:
		return false
	case
		reflect.Interface,
		reflect.Pointer,
		reflect.UnsafePointer,
		reflect.Slice,
		reflect.Map,
		reflect.Func,
		reflect.Chan:
		return v.IsNil()
	}
}

// IsBuiltIn returns true if t is a named built-in type.
func IsBuiltIn(t reflect.Type) bool {
	return t.PkgPath() == "" && t.Name() != ""
}

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
