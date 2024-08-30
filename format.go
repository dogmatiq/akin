package akin

import (
	"fmt"
	"io"
	"reflect"
	"slices"
	"strings"

	"github.com/dogmatiq/akin/internal/reflectx"
)

type formatter interface {
	formal() string
	human() string

	// hide returns the value represented as a type with the same underlying
	// type as the original value, but with no methods. This allows us to format
	// the value without recursion into the Format method.
	//
	// See https://github.com/golang/go/issues/51195#issuecomment-1563538796
	hide() any
}

func format(f formatter, s fmt.State, v rune) {
	if v != 's' {
		formatDefault(s, v, f)
	} else if s.Flag('+') {
		io.WriteString(s, f.human())
	} else {
		io.WriteString(s, f.formal())
	}
}

func formatDefault(f fmt.State, verb rune, v formatter) {
	s := fmt.FormatString(f, verb)
	h := v.hide()

	if _, ok := h.(fmt.Formatter); ok {
		panic("hide() must not return a value that implements fmt.Formatter")
	}

	fmt.Fprintf(f, s)
}

// sprintf behaves as per [fmt.sprintf], but renders [reflect.Type] and
// [reflect.Value] values using a more human-readable format.
func sprintf(format string, args ...any) string {
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

// sprint returns a human-readable representation of a value.
func sprint(v any) string {
	switch v := v.(type) {
	case reflect.Type:
		return formatT(v)
	case reflect.Value:
		return formatTV(v)
	case fmt.Stringer:
		return v.String()
	}
	return formatTV(reflectx.ValueOf(v))
}

// join renders a list of elements as a string, with each element
// separated by a separator string.
func join[T any](sep string, elements ...T) string {
	return join2(sep, sep, elements...)
}

// join2 renders a list of elements as a string, with each element
// separated by a separator string, and a different separator string
// before the last element.
func join2[T any](sep, lastSep string, elements ...T) string {
	var w strings.Builder
	last := len(elements) - 1

	for pos, elem := range elements {
		switch pos {
		case 0:
			// never add a separator before the first element, even if it's also
			// the last element.
		case last:
			w.WriteString(lastSep)
		default:
			w.WriteString(sep)
		}

		w.WriteString(sprint(elem))
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
	if reflectx.IsNil(v) {
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

func choose[B ~bool, T any](cond B, t, f T) T {
	if cond {
		return t
	}
	return f
}
