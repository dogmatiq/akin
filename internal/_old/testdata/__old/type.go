package akin

import (
	"reflect"
	"strings"
)

// HasType returns a [Predicate] that is satisfied when the [Value] has type T.
func HasType[T any]() TypeEquivalence {
	return TypeEquivalence{
		T: Type{
			rtype: reflect.TypeFor[T](),
		},
	}
}

// TypeEquivalence is a [Predicate] that is [True] when the [Value] has a
// specific [Type].
type TypeEquivalence struct {
	T Type
}

func (p TypeEquivalence) visitP(v PVisitor)        { v.TypeEquivalence(p) }
func (p TypeEquivalence) visitA(v PropertyVisitor) { v.TypeEquivalence(p) }
func (p TypeEquivalence) String() string           { return "𝒙 ⦂ " + p.T.String() }

func (e *evaluator) TypeEquivalence(p TypeEquivalence) {
	want := p.T
	got := e.Value.Type()

	e.Result = tern(got == want)
	e.Rationale = PropertyRationale{
		Property: p, // type equivalence is "self evident"
		Holds:    got == want,
	}
}

// Type represents a specific Go type.
type Type struct {
	rtype reflect.Type
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
