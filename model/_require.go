package akin

import (
	"fmt"
	"reflect"
)

func notEqualErr(got, want value) error {
	return fmt.Errorf("%T(%v) != %T(%v)", got.v, got.v, want.v, want.v)
}

func requireEqualTo(x, v value) error {
	if v.r.Equal(x.r) {
		return nil
	}
	return notEqualErr(v, x)
}

func requireType(t reflect.Type, v value) error {
	if v.dynamic == t {
		return nil
	}
	return fmt.Errorf("%s is not %s", v.dynamic, t)
}

func requireKind(k reflect.Kind, v value) error {
	if k == reflect.Invalid {
		panic("cannot require invalid kind")
	}

	if v.r.Kind() == k {
		return nil
	}

	if v.r.Kind() == reflect.Invalid {
		// TODO
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
		return fmt.Errorf("%s is not an %s", v.dynamic, k)
	}

	return fmt.Errorf("%s is not a %s", v.dynamic, k)
}

func requireConvertibleTo(t reflect.Type, v value) error {
	if !v.dynamic.ConvertibleTo(t) {
		return fmt.Errorf("cannot convert %s to %s", v.dynamic, t)
	}
	return nil
}
