package akin

import (
	"reflect"
	"strings"
)

// Type is the type of a Go [Value].
type Type struct{ rtype reflect.Type }

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

// TypeEq is a [Predicate] and [Attribute] that holds true when 𝒙 has a
// specific [Type].
//
// Type equality is notated using the set membership symbols, ∈ and ∉. For
// example, given 𝑷 ≔ ❨𝒙 ∈ int❩, then 𝑷❨𝒙❩ = 𝓽 if 𝒙 has a [Type] of [int].
type TypeEq struct {
	T Type
}

// VisitP calls the method on v associated with the predicate's type.
func (pa TypeEq) VisitP(v PVisitor) {
	v.TypeEq(pa)
}

// VisitA calls the method on v associated with the attribute's type.
func (pa TypeEq) VisitA(v AVisitor) {
	v.TypeEq(pa)
}

func (pa TypeEq) String() string {
	return stringP(pa)
}

func (s *identity) TypeEq(pa TypeEq) {
	s.fmt("𝒙 ∈ %s", pa.T)
}

func (s *inverted) TypeEq(pa TypeEq) {
	s.fmt("𝒙 ∉ %s", pa.T)
}

func (e *evaluator) TypeEq(pa TypeEq) {
	sameType := e.X.Type() == pa.T

	e.Px = truth(sameType)
	e.R = Ax{A: pa, Ax: sameType}
}
