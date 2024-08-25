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

func requireEquality(req, v reflect.Value) error {
	if v.Equal(req) {
		return nil
	}
	return notEqual(v, req)
}

func requireType(t reflect.Type, v reflect.Value) error {
	if t == v.Type() {
		return nil
	}
	return fmt.Errorf("%s is not %s", v.Type(), t)
}

func requireKind(k reflect.Kind, v reflect.Value) error {
	if k == reflect.Invalid {
		panic("cannot require invalid kind")
	}

	if v.Kind() == k {
		return nil
	}

	if v.Kind() == reflect.Invalid {
		return fmt.Errorf("invalid value")
	}

	switch k {
	case reflect.Array,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Interface,
		reflect.UnsafePointer:
		return fmt.Errorf("%s is not an %s", v.Type(), k)
	}

	return fmt.Errorf("%s is not a %s", v.Type(), k)
}
