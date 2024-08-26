package akin

import (
	"reflect"
)

// To returns the [Set] of values that are "akin to" the given model value.
func To(model any) Set {
	v := valueOf(model)
	return fromModel(v)
}

func fromModel(v reflect.Value) Set {
	if v.Type().PkgPath() == "" {
		switch v.Kind() {
		case reflect.Array:
		case reflect.Chan:
		case reflect.Func:

		case reflect.Interface:
			if v.Type() == reflect.TypeFor[any]() && v.IsNil() {
				return Nil
			}

		case reflect.Map:
		case reflect.Pointer:
		case reflect.Slice:
		case reflect.String:
		case reflect.Struct:
		}
	}

	return singleton{v}
}
