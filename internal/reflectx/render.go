package reflectx

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

// Sprintf behaves as per [fmt.Sprintf], but renders [reflect.Type] and
// [reflect.Value] values using a more human-readable format.
func Sprintf(format string, args ...any) string {
	cloned := false
	set := func(i int, v any) {
		if !cloned {
			cloned = true
			args = slices.Clone(args)
		}
		args[i] = v
	}

	for i, v := range args {
		switch v := v.(type) {
		case reflect.Type:
			set(i, formatT(v))
		case reflect.Value:
			set(i, formatTV(v))
		}
	}

	return fmt.Sprintf(format, args...)
}

// Sprint returns a human-readable representation of a value.
func Sprint(v any) string {
	switch v := v.(type) {
	case reflect.Type:
		return formatT(v)
	case reflect.Value:
		return formatTV(v)
	case fmt.Stringer:
		return v.String()
	}
	return formatTV(ValueOf(v))
}

// SprintList renders a list of elements as a string, with each element
// separated by a separator string.
func SprintList[T any](sep string, elements ...T) string {
	var w strings.Builder

	for i, e := range elements {
		if i > 0 {
			w.WriteString(sep)
		}
		w.WriteString(Sprint(e))
	}

	return w.String()
}

func formatTV(v reflect.Value) string {
	t := v.Type()

	rv := formatV(v)

	switch t {
	case reflect.TypeOf(""),
		reflect.TypeOf(true),
		reflect.TypeOf(0),
		reflect.TypeOf(0.0),
		reflect.TypeOf(0 + 0i):
		return rv
	}

	rt := formatT(t)

	return fmt.Sprintf("%s(%s)", rt, rv)
}

func formatT(t reflect.Type) string {
	if t == reflect.TypeFor[any]() {
		return "any"
	}

	s := strings.ReplaceAll(
		t.String(),
		" {",
		"{",
	)

	if strings.ContainsAny(s, " *({") {
		s = "(" + s + ")"
	}

	return s
}

func formatV(v reflect.Value) string {
	if IsNil(v) {
		return "nil"
	}

	if v.CanComplex() {
		s := fmt.Sprint(v.Interface())
		return s[1 : len(s)-1]
	}

	if v.CanFloat() {
		s := fmt.Sprint(v.Interface())
		if strings.ContainsAny(s, ".eE") {
			return s
		}
		return s + ".0"
	}

	if v.Kind() == reflect.String {
		return fmt.Sprintf("%q", v.Interface())
	}

	return fmt.Sprintf("%v", v.Interface())
}
