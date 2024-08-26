package akin

import (
	"fmt"
	"reflect"
)

func renderV(v reflect.Value) string {
	if v.Kind() == reflect.Invalid {
		return "any(nil)"
	}

	x := v.Interface()
	return fmt.Sprintf("%T(%v)", x, x)
}

func renderT(t reflect.Type) string {
	return t.String() // TODO
}
