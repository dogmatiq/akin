package akin

import "fmt"

type (
	// A Predicate describes some criteria that a [Value] may or may not
	// satisfy.
	//
	// Predicates are represented by a "mathematical bold italic capital"
	// letter, typically 𝑷, although other letters may be used when discussing
	// multiple predicates.
	//
	// Similarly, values are represented by a "mathematical bold italic small"
	// letter, typically 𝒙, although other letters may be used when discussing
	// multiple values.
	Predicate interface {
		fmt.Stringer

		acceptPredicateVisitor(PredicateVisitor)
	}

	// An Assertion is a kind of [Predicate] that can directly test if a [Value]
	// satisfies the predicate's criteria.
	//
	// To determine if 𝒙 satisfies 𝑷, we "evaluate 𝑷 of 𝒙", written 𝑷❨𝒙❩,
	// using the [Eval] function.
	Assertion interface {
		Predicate

		acceptAssertionVisitor(AssertionVisitor)
	}
)

// A PredicateVisitor is an algorithm that applies different logic for each
// [Predicate] type.
type PredicateVisitor interface {
	AssertionVisitor

	VisitTypeEq(TypeEq)
	VisitValueEq(ValueEq)
}

// An AssertionVisitor is an algorithm that applies different logic for each
// [Assertion] type.
type AssertionVisitor interface {
	VisitConst(Const)
	VisitNilness(Nilness)
	VisitTypehood(Typehood)
}
