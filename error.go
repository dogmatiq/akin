package akin

import (
	"fmt"
	"reflect"
)

func notEqual(a, b reflect.Value) error {
	var va, vb any

	if a.Kind() != reflect.Invalid {
		va = a.Interface()
	}

	if b.Kind() != reflect.Invalid {
		vb = b.Interface()
	}

	return fmt.Errorf("%v != %v", va, vb)
}
