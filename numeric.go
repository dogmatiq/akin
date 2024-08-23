package akin

import (
	"fmt"
	"reflect"
)

// numeric returns a predicate that checks if it's argument can be losslessly
// converted to the same type as the spec value, and when converted the values
// are equal.
func numeric(spec reflect.Value) func(reflect.Value) error {
	specT := spec.Type()
	isNeg := isNegative(spec)

	return func(v reflect.Value) error {
		t := v.Type()

		if !t.ConvertibleTo(specT) {
			return fmt.Errorf("cannot convert %s to %s", t, specT)
		}

		if !specT.ConvertibleTo(t) {
			return fmt.Errorf("cannot convert back to %s from %s", t, specT)
		}

		if !v.CanConvert(specT) {
			return fmt.Errorf("cannot convert %#v value to %s", v.Interface(), specT)
		}

		if isNeg != isNegative(v) {
			return notEqual(v, spec)
		}

		converted := v.Convert(specT)
		if !converted.Equal(spec) {
			return notEqual(v, spec)
		}

		reverted := converted.Convert(v.Type())
		if !reverted.Equal(v) {
			return fmt.Errorf(
				"type conversion is lossy, %T(%v) != %T(%v)",
				reverted.Interface(),
				reverted.Interface(),
				v.Interface(),
				v.Interface(),
			)
		}

		return nil
	}
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
