package akin

import "reflect"

// Spec describes a value against which other values can be compared for
// likeness.
type Spec interface {
	Compare(any) ComparisonResult
}

// ComparisonResult encapsulates the result of comparing a value to a [Spec].
type ComparisonResult struct {
	Error error
}

// NewSpec returns a new specification that checks if values are "like" v.
func NewSpec(v any) Spec {
	return compile(reflect.ValueOf(v))
}

// Test returns an error if v is not "like" the given specification value.
//
// If multiple values will be compared against the same specification, it is
// more efficient to build the specification once using [NewSpec].
func Test(spec, v any) error {
	return NewSpec(spec).Compare(v).Error
}

func compile(spec reflect.Value) Spec {
	switch spec.Kind() {
	case reflect.Bool:
		return compileBool(spec)

	case
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return compileNumeric(spec)

	case reflect.Slice:
		return compileSlice(spec)

	case reflect.Map:
		return compileMap(spec)

	case reflect.Func:
		return compileFunc(spec)

		// case reflect.Uintptr:
		// case reflect.Complex64:
		// case reflect.Complex128:
		// case reflect.Array:
		// case reflect.Chan:
		// case reflect.Interface:
		// case reflect.Pointer:
		// case reflect.String:
		// case reflect.Struct:
		// case reflect.UnsafePointer:
	}

	return specFunc(func(arg reflect.Value) error {
		if !arg.Equal(spec) {
			return notEqual(arg, spec)
		}
		return nil
	})
}

type specFunc func(reflect.Value) error

func (s specFunc) Compare(arg any) ComparisonResult {
	err := s(reflect.ValueOf(arg))
	return ComparisonResult{Error: err}
}
