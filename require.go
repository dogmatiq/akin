package akin

import (
	"fmt"
	"reflect"
)

func notEqualErr(got, want Value) error {
	return fmt.Errorf("%T(%v) != %T(%v)", got.value, got.value, want.value, want.value)
}

func requireEqualTo(x, v Value) error {
	if v.rvalue.Equal(x.rvalue) {
		return nil
	}
	return notEqualErr(v, x)
}

func requireType(t reflect.Type, v Value) error {
	if v.rtype == t {
		return nil
	}
	return fmt.Errorf("%s is not %s", v.rtype, t)
}

func requireKind(k reflect.Kind, v Value) error {
	if k == reflect.Invalid {
		panic("cannot require invalid kind")
	}

	if v.rvalue.Kind() == k {
		return nil
	}

	if v.rvalue.Kind() == reflect.Invalid {
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
		return fmt.Errorf("%s is not an %s", v.rtype, k)
	}

	return fmt.Errorf("%s is not a %s", v.rtype, k)
}

func requireConvertibleTo(t reflect.Type, v Value) error {
	if !v.rtype.ConvertibleTo(t) {
		return fmt.Errorf("cannot convert %s to %s", v.rtype, t)
	}
	return nil
}
