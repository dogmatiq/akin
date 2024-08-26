package akin

import "reflect"

// renderType returns a human-readable name for t.
func renderType(t reflect.Type) string {
	// TODO
	return t.String()
}

func renderValue(v reflect.Value) string {
	// TODO
	return v.String()
}
