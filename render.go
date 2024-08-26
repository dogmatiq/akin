package akin

import "reflect"

// renderType returns a human-readable name for t.
func renderType(t reflect.Type) string {
	// TODO
	return t.String()
}

func renderValue(v reflect.Value) string {
	if v.Kind() == reflect.Invalid {
		return "any(nil)"
	}

	// TODO
	return v.String()
}
