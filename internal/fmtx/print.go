package fmtx

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dogmatiq/akin/internal/reflectx"
	"github.com/dogmatiq/akin/internal/slicex"
)

// F is a specialization of [fmt.F] that renders [reflect.Type] and
// [reflect.Value] values using more concise and readable representations.
func F(format string, args ...any) string {
	return fmt.Sprintf(
		format,
		slicex.Map(
			args,
			func(arg any) (any, bool) {
				switch arg := arg.(type) {
				case reflect.Type:
					return formattedType{arg}, true
				case reflect.Value:
					return formattedValue{arg}, true
				}
				return arg, false
			},
		)...,
	)
}

// P is a specialization of [fmt.P] that renders [reflect.Type] and
// [reflect.Value] values using more concise and readable representations.
func P(v any) string {
	switch v := v.(type) {
	case reflect.Type:
		return formattedType{v}.String()
	case reflect.Value:
		return formattedValue{v}.String()
	default:
		return fmt.Sprint(v)
	}
}

type (
	formattedType  struct{ t reflect.Type }
	formattedValue struct{ v reflect.Value }
)

func (f formattedType) String() string {
	if f.t == reflect.TypeFor[any]() {
		return "any"
	}

	s := strings.ReplaceAll(
		f.t.String(),
		" {",
		"{",
	)

	if strings.ContainsAny(s, " *({") {
		s = "(" + s + ")"
	}

	return s
}

func (f formattedValue) String() string {
	switch f.v.Type() {
	case reflect.TypeOf(""),
		reflect.TypeOf(true),
		reflect.TypeOf(0),
		reflect.TypeOf(0.0),
		reflect.TypeOf(0 + 0i):
		return f.valuePart()
	}

	return f.typePart() + "(" + f.valuePart() + ")"
}

func (f formattedValue) typePart() string {
	return formattedType{f.v.Type()}.String()
}

func (f formattedValue) valuePart() string {
	if reflectx.IsNil(f.v) {
		return "nil"
	}

	if f.v.CanComplex() {
		s := fmt.Sprint(f.v.Interface())
		return s[1 : len(s)-1]
	}

	if f.v.CanFloat() {
		s := fmt.Sprint(f.v.Interface())
		if strings.ContainsAny(s, ".eE") {
			return s
		}
		return s + ".0"
	}

	if f.v.Kind() == reflect.String {
		return fmt.Sprintf("%q", f.v.Interface())
	}

	return fmt.Sprintf("%v", f.v.Interface())
}
