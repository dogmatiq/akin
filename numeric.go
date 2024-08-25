package akin

import (
	"fmt"
	"reflect"
)

func compileNumeric(spec reflect.Value) Spec {
	specT := spec.Type()
	isNeg := isNegative(spec)

	return specFunc(func(arg reflect.Value) error {
		argT := arg.Type()

		if !argT.ConvertibleTo(specT) {
			return fmt.Errorf("cannot convert %s to %s", argT, specT)
		}

		if !specT.ConvertibleTo(argT) {
			return fmt.Errorf("cannot convert back to %s from %s", argT, specT)
		}

		if !arg.CanConvert(specT) {
			return fmt.Errorf("cannot convert %#v value to %s", arg.Interface(), specT)
		}

		if isNeg != isNegative(arg) {
			return notEqual(arg, spec)
		}

		converted := arg.Convert(specT)
		if !converted.Equal(spec) {
			return notEqual(arg, spec)
		}

		reverted := converted.Convert(arg.Type())
		if !reverted.Equal(arg) {
			return fmt.Errorf(
				"type conversion is lossy, %T(%v) != %T(%v)",
				reverted.Interface(),
				reverted.Interface(),
				arg.Interface(),
				arg.Interface(),
			)
		}

		return nil
	})
}

func isNegative(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() < 0
	case reflect.Float32, reflect.Float64:
		return v.Float() < 0
	}
	return false
}
