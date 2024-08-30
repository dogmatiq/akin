package akin

import "fmt"

// A Reason is a justification for the outcome of an [Evaluation].
type Reason interface {
	visitReason(ReasonVisitor)
}

type isReason[T interface {
	fmt.Stringer
	Reason
}] struct{}

var (
	_ = isReason[PredicateIsConstant]{}
	_ = isReason[NoConstituents]{}
	_ = isReason[ConstituentSatisfied]{}
	_ = isReason[PropertySatisfied]{}
	_ = isReason[PropertyViolated]{}
)

// A ReasonVisitor encapsulates logic specific to each [Reason] type.
type ReasonVisitor interface {
	VisitPredicateIsConstantReason(PredicateIsConstant)
	VisitNoConstituentsReason(NoConstituents)
	VisitConstituentSatisfiedReason(ConstituentSatisfied)
	VisitAllConstituentsViolatedReason(AllConstituentsViolated)
	VisitPropertySatisfiedReason(PropertySatisfied)
	VisitPropertyViolatedReason(PropertyViolated)
}

// PredicateIsConstant is a [Reason] indicating that the predicate produces the
// same result for all values.
type PredicateIsConstant struct {
	P Predicate
}

func (r PredicateIsConstant) visitReason(v ReasonVisitor) {
	v.VisitPredicateIsConstantReason(r)
}

func (r PredicateIsConstant) String() string {
	return r.P.human()
}

// NoConstituents is a [Reason] indicating that a predicate has no constituent
// predicates.
type NoConstituents struct {
	P Predicate
}

func (r NoConstituents) visitReason(v ReasonVisitor) {
	v.VisitNoConstituentsReason(r)
}

func (r NoConstituents) String() string {
	return "𝑷 has no constituent predicates"
}

// ConstituentSatisfied is a [Reason] indicating that one of the
// constituent predicates is satisfied.
type ConstituentSatisfied struct {
	Evaluation Evaluation
}

func (r ConstituentSatisfied) visitReason(v ReasonVisitor) {
	v.VisitConstituentSatisfiedReason(r)
}

func (r ConstituentSatisfied) String() string {
	return fmt.Sprintf(
		"constituent predicate %s is satisfied, because %s",
		r.Evaluation.P,
		r.Evaluation.Reason,
	)
}

// AllConstituentsViolated is a [Reason] indicating that all of the
// constituent predicates are violated.
type AllConstituentsViolated struct {
	P Predicate
}

func (r AllConstituentsViolated) visitReason(v ReasonVisitor) {
	v.VisitAllConstituentsViolatedReason(r)
}

func (r AllConstituentsViolated) String() string {
	return "all constituent predicates are violated"
}

// PropertySatisfied is a [Reason] indicating that [Property] holds true for 𝑥.
type PropertySatisfied struct {
	P Property
}

func (r PropertySatisfied) visitReason(v ReasonVisitor) {
	v.VisitPropertySatisfiedReason(r)
}

func (r PropertySatisfied) String() string {
	return r.P.human()
}

// PropertyViolated is a [Reason] indicating that a [Property] does not hold
// true for 𝑥.
type PropertyViolated struct {
	P Property
}

func (r PropertyViolated) visitReason(v ReasonVisitor) {
	v.VisitPropertyViolatedReason(r)
}

func (r PropertyViolated) String() string {
	return r.P.inverse()
}
