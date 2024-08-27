package akin

import (
	"fmt"
	"reflect"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// A Predicate describes a condition that a value must satisfy.
type Predicate interface {
	fmt.Stringer

	// Eval evaluates v against the predicate.
	Eval(v any) Evaluation

	// Simplify returns the simplest possible predicate that is equivalent to
	// this predicate.
	//
	// p is always a non-nil predicate, even if it is the same predicate.
	//
	// simplified is true if the predicate was simplified, or false if the same
	// predicate was returned.
	Simplify() (p Predicate, simplified bool)
}

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
			return AnyOf(IsNil, EqualTo(uintptr(0)))
		}

	case reflect.Map:
	case reflect.Pointer:
	case reflect.Slice:
	case reflect.String:
	case reflect.Struct:
	}

	return convertibleTo{v}
}
