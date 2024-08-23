package akin

import (
	"fmt"
	"reflect"
)

func notEqual(a, b reflect.Value) error {
	return fmt.Errorf("%v != %v", a.Interface(), b.Interface())
}
