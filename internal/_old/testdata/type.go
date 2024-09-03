package akin

import (
	"fmt"
	"reflect"
)

// HasType returns a [Predicate] that is satisfied when 𝒙 has type T.
func HasType[T any]() TypeEquivalence {
	return TypeEquivalence{reflect.TypeFor[T]()}
}

// TypeEquivalence is a [Predicate] and a [Property] that is satisfied when 𝒙
// has a specific type.
type TypeEquivalence struct {
	T reflect.Type
}

// Format implements the [fmt.Formatter] interface.
func (p TypeEquivalence) Format(s fmt.State, v rune) {
	format(p, s, v)
}
func (p TypeEquivalence) hide() any {
	type T = TypeEquivalence
	type Type T
	return Type(p)
}

func (p TypeEquivalence) formal() string {
	return renderf("𝒙 ⦂ %s", p.T)
}

func (p TypeEquivalence) human() string {
	return renderf("𝒙 has type %s", p.T)
}

func (p TypeEquivalence) inverse() string {
	return renderf("𝒙 does not have type %s", p.T)
}

func (p TypeEquivalence) visitPredicate(v PredicateVisitor) {
	v.VisitTypeEquivalencePredicate(p)
}

func (p TypeEquivalence) visitProperty(v PropertyVisitor) {
	v.VisitTypeEquivalenceProperty(p)
}

func (i *inverter) VisitTypeEquivalencePredicate(p TypeEquivalence) {
	i.Q = Not(p)
}

func (r *reducer) VisitTypeEquivalencePredicate(p TypeEquivalence) {
	r.Q = p
}

func (e *evaluator) VisitTypeEquivalencePredicate(p TypeEquivalence) {
	e.SetProperty(e.T == p.T, p)
}
