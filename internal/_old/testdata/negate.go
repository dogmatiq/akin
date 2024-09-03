package akin

import "fmt"

// Not returns a [Predicate] 𝑷 that is satisfied when 𝒙 does not satisfy 𝐐.
func Not(q Predicate) Predicate {
	return Negation{q}
}

// Negation is a [Predicate] that is satisfied when 𝒙 does not satisfy its
// constituent predicate 𝐐.
type Negation struct {
	Q Predicate
}

// Format implements the [fmt.Formatter] interface.
func (p Negation) Format(f fmt.State, v rune) {
	format(p, f, v)
}

func (p Negation) hide() any {
	type T = Negation
	type Negation T
	return T(p)
}

func (p Negation) formal() string {
	return "¬ " + parens(p.Q.formal())
}

func (p Negation) human() string {
	return renderf("𝒙 does not satisfy %s", p.Q)
}

func (p Negation) visitPredicate(v PredicateVisitor) {
	v.VisitNegatePredicate(p)
}

func (i *inverter) VisitNegatePredicate(p Negation) {
	i.Q = p.Q
}

func (r *reducer) VisitNegatePredicate(p Negation) {
	// TODO: collapse double negations.
	r.Q = p
}

func (e *evaluator) VisitNegatePredicate(p Negation) {
	qe := eval(p.Q, e.X)

	if qe.IsSatisfied {
		e.IsSatisfied = false
		e.Reason = ConstituentSatisfied{qe}
	} else {
		e.IsSatisfied = true
		e.Reason = ConstituentViolated{qe}
	}
}
