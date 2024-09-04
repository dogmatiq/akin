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
