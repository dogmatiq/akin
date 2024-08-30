package akin

import (
	"fmt"

	"github.com/dogmatiq/akin/internal/reflectx"
)

const (
	// IsNil is a [Predicate] that is satisfied when 𝑥 is nil, regardless of
	// its type.
	IsNil Nilness = true

	// IsNonNil is a [Predicate] that is satisfied when 𝑥 is non-nil,
	// regardless of its type.
	IsNonNil Nilness = false
)

// Nilness is a [Predicate] and [Property] that is satisfied when 𝑥 has
// "nilness" equal to the predicate's value.
type Nilness bool

// Format implements the [fmt.Formatter] interface.
func (p Nilness) Format(f fmt.State, verb rune) {
	format(p, f, verb)
}

func (p Nilness) hide() any {
	type T = Nilness
	type Nilness T
	return Nilness(p)
}

func (p Nilness) formal() string {
	return choose(p, "𝑥 = nil", "𝑥 ≠ nil")
}

func (p Nilness) human() string {
	return choose(p, "𝑥 is nil", "𝑥 is not nil")
}

func (p Nilness) inverse() string {
	return (!p).human()
}

func (p Nilness) visitPredicate(v PredicateVisitor) {
	v.VisitNilnessPredicate(p)
}

func (e *evaluator) VisitNilnessPredicate(p Nilness) {
	e.IsSatisfied = reflectx.IsNil(e.X) == bool(p)

	prop := ValueEquivalence{"nil"}
	if e.IsSatisfied && !bool(p) {
		e.Reason = PropertySatisfied{prop}
	} else {
		e.Reason = PropertyViolated{prop}
	}
}
