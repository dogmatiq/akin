package akin

import (
	"fmt"
	"reflect"
)

// New returns a predicate function that tests if its argument is akin to the
// spec value.
func New[T any](spec T) func(T) error {
	pred := newPredicate(reflect.ValueOf(spec))
	return func(v T) error {
		return pred(reflect.ValueOf(v))
	}
}

func newPredicate(spec reflect.Value) func(reflect.Value) error {
	if spec.Kind() == reflect.Invalid {
		return func(v reflect.Value) error {
			if v.IsValid() {
				return fmt.Errorf("expected nil, got %s", v.Type())
			}
			return nil
		}
	}

	switch spec.Kind() {
	case
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:

		return numeric(spec)
	}

	// case reflect.Bool:
	// case reflect.Uintptr:
	// case reflect.Complex64:
	// case reflect.Complex128:
	// case reflect.Array:
	// case reflect.Chan:
	// case reflect.Func:
	// case reflect.Interface:
	// case reflect.Map:
	// case reflect.Pointer:
	// case reflect.Slice:
	// case reflect.String:
	// case reflect.Struct:
	// case reflect.UnsafePointer:

	return func(v reflect.Value) error {
		if !spec.Equal(v) {
			return notEqual(v, spec)
		}
		return nil
	}
}

// Test returns true if the candidate value v is akin to the spec value.
func Test(spec, v any) error {
	return New(spec)(v)
}
