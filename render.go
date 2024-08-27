package akin

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dogmatiq/akin/internal/reflectx"
)

func render(v reflect.Value) string {
	t := v.Type()

	rv := renderV(v)

	switch t {
	case reflect.TypeFor[string](),
		reflect.TypeFor[bool]():
		return rv
	}

	rt := renderT(t)
	if strings.ContainsAny(rt, " *({") {
		rt = "(" + rt + ")"
	}

	return fmt.Sprintf("%s(%s)", rt, rv)
}

func renderT(t reflect.Type) string {
	if t == reflect.TypeFor[any]() {
		return "any"
	}

	return strings.ReplaceAll(
		t.String(),
		" {",
		"{",
	)
}

func renderV(v reflect.Value) string {
	if v.Kind() == reflect.String {
		return fmt.Sprintf("%q", v.Interface())
	}

	if reflectx.IsNil(v) {
		return "nil"
	}

	return fmt.Sprintf("%v", v.Interface())
}
