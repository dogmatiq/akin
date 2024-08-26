package akin

import (
	"fmt"
	"reflect"
	"strings"
)

func renderV(v reflect.Value) string {
	rt := renderT(v.Type())

	if strings.ContainsAny(rt, "*[({") {
		rt = "(" + rt + ")"
	}

	rv := "nil"
	if !isNil(v) {
		rv = fmt.Sprintf("%v", v.Interface())
	}

	return fmt.Sprintf(
		"%s(%s)",
		renderT(v.Type()),
		rv,
	)
}

func renderT(t reflect.Type) string {
	if t == reflect.TypeFor[any]() {
		return "any"
	}
	return strings.ReplaceAll(
		t.String(),
		" ",
		"",
	)
}
