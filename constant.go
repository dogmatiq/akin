package akin

import "fmt"

const (
	// Top is a [Predicate] that is satisfied by any value.
	Top Constant = true

	// Bottom is a [Predicate] that is not satisfied by any values.
	Bottom Constant = false
)

// Constant is a [Predicate] that is either always satisfied or never satisfied,
// regardless of any property of 𝑥.
//
// The predicate that is always satisfied is called [Top], and is denoted by the
// symbol "⊤". The predicate that is never satisfied is called [Bottom], and is
// denoted by the symbol "⊥".
type Constant bool

// Format implements the [fmt.Formatter] interface.
func (p Constant) Format(s fmt.State, v rune) {
	format(p, s, v)
}

func (p Constant) hide() any {
	type T = Constant
	type Constant T
	return Constant(p)
}

func (p Constant) formal() string {
	return choose(p, "⊤", "⊥")
}

func (p Constant) human() string {
	return "𝑷 is " + choose(p, "satisfied", "violated") + " by all values"
}

func (p Constant) visitPredicate(v PredicateVisitor) {
	v.VisitConstantPredicate(p)
}

func (e *evaluator) VisitConstantPredicate(p Constant) {
	e.IsSatisfied = bool(p)
	e.Reason = PredicateIsConstant{p}
}
