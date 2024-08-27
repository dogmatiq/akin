package akin

import (
	"reflect"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// To returns a [Predicate] that matches values that are "akin to" the given
// model value.
func To(model any) Predicate {
	v := reflectx.ValueOf(model)
	return fromModel(v)
}

func fromModel(v reflect.Value) Predicate {
	if v.Type().PkgPath() != "" {
		return equalTo{v}
	}

	switch v.Kind() {
	case reflect.Array:
	case reflect.Chan:
	case reflect.Func:

	case reflect.Interface:
		if v.Type() == reflect.TypeFor[any]() && v.IsNil() {
			return Or(IsNil, EqualTo(uintptr(0)))
		}

	case reflect.Map:
	case reflect.Pointer:
	case reflect.Slice:
	case reflect.String:
	case reflect.Struct:
	}

	return convertibleTo{v}
}
