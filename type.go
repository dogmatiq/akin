package akin

import (
	"reflect"
	"strings"
)

// Type is the type of a Go [Value].
//
// The type of a value is notated using 𝝉 (mathematical bold italic small tau).
// For example, 𝝉❨𝒙❩ refers to the [Type] of the 𝒙 value.
type Type struct{ ref reflect.Type }

func typeFor[T any]() Type {
	return Type{reflect.TypeFor[T]()}
}

// isNilable returns true if a value of type t can be nil.
func (t Type) isNilable() bool {
	switch t.ref.Kind() {
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
	return t.ref.PkgPath() == "" && t.ref.Name() != ""
}

// isDefaultForConstant if t is the default type used to represent an untyped
// constant.
//
// See "default type" under https://go.dev/ref/spec#Constants
func (t Type) isDefaultForConstant() bool {
	switch t.ref.Kind() {
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
	if t.ref == reflect.TypeFor[any]() {
		return "any"
	}

	s := strings.ReplaceAll(
		t.ref.String(),
		" {",
		"{",
	)

	if strings.ContainsAny(s, " *({") {
		s = "(" + s + ")"
	}

	return s
}
