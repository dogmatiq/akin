package akin

import "reflect"

// Spec describes a value against which other values can be compared for
// likeness.
type Spec interface {
	Compare(Value) ComparisonResult
}

// ComparisonResult encapsulates the result of comparing a value to a [Spec].
type ComparisonResult struct {
	Error error
}

// Compile returns a new specification that checks if values are "like" v.
func Compile(v any) Spec {
	return compile(ValueOf(v))
}

// Test returns an error if v is not "like" the given specification value.
//
// If multiple values will be compared against the same specification, it is
// more efficient to build the specification once using [Compile].
func Test(spec, v any) error {
	return Compile(spec).
		Compare(ValueOf(v)).
		Error
}

func compile(spec Value) Spec {
	if spec.value != nil {
		switch spec.rvalue.Kind() {
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
	}

	return specFunc(func(v Value) error {
		return requireEqualTo(spec, v)
	})
}

type specFunc func(Value) error

func (s specFunc) Compare(v Value) ComparisonResult {
	err := s(v)
	return ComparisonResult{Error: err}
}
