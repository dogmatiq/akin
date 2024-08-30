package akin

import (
	"fmt"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// A Predicate 𝑷 describes a condition or set of conditions that a value 𝑥
// either satisfies or violates.
//
// Many [Predicate] types also implement [Property].
//
// All predicates can be formatted with the [fmt] package. The "%s" verb
// produces a (somewhat) formal representation of the predicate, whereas "%+s"
// produces a more verbose human-readable description.
type Predicate interface {
	formatter

	visitPredicate(PredicateVisitor)
}

// Eval evaluates a value against a predicate, that is 𝑷❨𝑥❩.
func Eval(p Predicate, x any) Evaluation {
	return eval(p, reflectx.ValueOf(x))
}

// Reduce returns the simplest form of p.
func Reduce(p Predicate) Predicate {
	return reduce(p)
}

// Same returns true if a and b are the "same" predicate.
//
// Two predicates are the same if they are the same type and have equivalent
// parameters. This is a kind of weak equality that respects the commutative
// properties of some predicate types.
func Same(a, b Predicate) bool {
	return same(a, b)
}

type isPredicate[T interface {
	fmt.Formatter
	Predicate
}] struct{}

var (
	_ = isPredicate[Constant]{}
	_ = isPredicate[Equal]{}
	_ = isPredicate[Equivalence]{}
	_ = isPredicate[Nilness]{}
	_ = isPredicate[Or]{}
	_ = isPredicate[TypeEquivalence]{}
)

// A PredicateVisitor encapsulates logic specific to each [Predicate] type.
type PredicateVisitor interface {
	VisitConstantPredicate(Constant)
	VisitEqualPredicate(Equal)
	VisitEquivalentPredicate(Equivalence)
	VisitNilnessPredicate(Nilness)
	VisitOrPredicate(Or)
	VisitTypeEquivalencePredicate(TypeEquivalence)
}

// VisitPredicate calls the appropriate method on v for the given predicate.
func VisitPredicate(p Predicate, v PredicateVisitor) {
	p.visitPredicate(v)
}
