package akin

import (
	"reflect"
	"strings"
)

// Type is the type of a Go [Value].
type Type struct{ rtype reflect.Type }

func typeFor[T any]() Type {
	return Type{reflect.TypeFor[T]()}
}

// isNilable returns true if a value of type t can be nil.
func (t Type) isNilable() bool {
	switch t.rtype.Kind() {
	default:
		return false
	case
		reflect.Chan,
		reflect.Func,
		reflect.Interface,
		reflect.Map,
		reflect.Pointer,
		reflect.Slice,
		reflect.UnsafePointer:
		return true
	}
}

// isBuiltIn returns true if t is a named built-in type.
func (t Type) isBuiltIn() bool {
	return t.rtype.PkgPath() == "" && t.rtype.Name() != ""
}

// isDefaultForConstant if t is the default type used to represent an untyped
// constant.
//
// See "default type" under https://go.dev/ref/spec#Constants
func (t Type) isDefaultForConstant() bool {
	switch t.rtype.Kind() {
	case reflect.Bool,
		reflect.Int32, // rune
		reflect.Int,
		reflect.Float64,
		reflect.Complex128,
		reflect.String:
		return t.isBuiltIn()
	default:
		return false
	}
}

func (t Type) String() string {
	if t.rtype == reflect.TypeFor[any]() {
		return "any"
	}

	s := strings.ReplaceAll(
		t.rtype.String(),
		" {",
		"{",
	)

	if strings.ContainsAny(s, " *({") {
		s = "(" + s + ")"
	}

	return s
}
