package akin

import (
	"reflect"
)

// Compile builds a predicate function that tests if its argument is akin to the
// given specification value.
func Compile[T any](spec T) func(T) error {
	pred := compile(reflect.ValueOf(spec))
	return func(arg T) error {
		return pred(reflect.ValueOf(arg))
	}
}

// Test returns true if the candidate value v is akin to the spec value.
func Test(spec, v any) error {
	return Compile(spec)(v)
}

func compile(spec reflect.Value) func(reflect.Value) error {
	switch spec.Kind() {
	case
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return compileNumeric(spec)

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
	}

	return func(arg reflect.Value) error {
		if !arg.Equal(spec) {
			return notEqual(arg, spec)
		}
		return nil
	}
}
