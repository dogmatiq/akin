package akin

import (
	"fmt"
	"reflect"
)

// valueOf returns the reflect.Value of v, with special handling for invalid
// values (i.e, nil interface values).
func valueOf(v any) reflect.Value {
	if r := reflect.ValueOf(v); r.IsValid() {
		return r
	}
	if r := reflect.ValueOf(&v).Elem(); r.IsValid() {
		return r
	}
	panic(fmt.Sprintf("cannot reflect on value of %v", v))
}
