package akin

import (
	"fmt"
)

// A Rationale describes the logical reasoning that justifies an [Evaluation].
//
// The ∵ symbol (because) is used to represent a rationale.
type Rationale interface {
	fmt.Stringer
	acceptRationaleVisitor(RationaleVisitor)
}

type (
	// ConstRationale is the [Rationale] provided when a [Predicate] always
	// produces the same [Result] regardless of the [Value].
	ConstRationale struct {
		Predicate     Predicate
		PredicateExpr Expr
	}

	// VacuousRationale is the [Rationale] provided when a [Predicate] makes no
	// meaningful assertion about the [Value].
	//
	// For example, a compound predicate with no constituent predicates is
	// considered vacuous. The [Result] of a vacuous predicate is always
	// [Undefined].
	VacuousRationale struct {
		Predicate     Predicate
		PredicateExpr Expr
	}

	// AssertionRationale is the [Rationale] provided when a [Result] is
	// determined by evaluation of an [Assertion].
	AssertionRationale struct {
		Assertion     Assertion
		AssertionExpr Expr
		Value         Value
		ValueExpr     Expr
		Result        Result
		Rationale     Rationale
	}

	// IntrinsicRationale is the [Rationale] provided when a [Result] is derived
	// from a [Predicate] that is not an [Assertion].
	//
	// Such predicates describe intrinsic attributes of a value which may be
	// abstract in nature.
	IntrinsicRationale struct {
		Predicate     Predicate
		PredicateExpr Expr
		Value         Value
		ValueExpr     Expr
		Result        bool
	}
)

// RationaleVisitor is an algorithm with logic specific to each [Rationale] type.
type RationaleVisitor interface {
	VisitConstRationale(ConstRationale)
	VisitVacuousRationale(VacuousRationale)
	VisitAssertionRationale(AssertionRationale)
	VisitIntrinsicRationale(IntrinsicRationale)
}

func (r ConstRationale) acceptRationaleVisitor(v RationaleVisitor)     { v.VisitConstRationale(r) }
func (r VacuousRationale) acceptRationaleVisitor(v RationaleVisitor)   { v.VisitVacuousRationale(r) }
func (r AssertionRationale) acceptRationaleVisitor(v RationaleVisitor) { v.VisitAssertionRationale(r) }
func (r IntrinsicRationale) acceptRationaleVisitor(v RationaleVisitor) { v.VisitIntrinsicRationale(r) }

func (r ConstRationale) String() string     { return rationaleToString(r) }
func (r VacuousRationale) String() string   { return rationaleToString(r) }
func (r AssertionRationale) String() string { return rationaleToString(r) }
func (r IntrinsicRationale) String() string { return rationaleToString(r) }

func (rr *rationaleRenderer) VisitConstRationale(r ConstRationale) {
	rr.Render("%s is constant", r.PredicateExpr)
}

func (rr *rationaleRenderer) VisitVacuousRationale(r VacuousRationale) {
	rr.Render("%s is vacuous", r.PredicateExpr)
}

func (rr *rationaleRenderer) VisitAssertionRationale(r AssertionRationale) {
	pr := &predicateRenderer{
		PredicateExpr: r.AssertionExpr,
		ValueExpr:     r.ValueExpr,
		Form:          affirmativeForm,
		Parenthesize:  true,
		Output:        rr.Output,
	}

	rr.Render("%s ≔ %s, %s ≔ ", r.ValueExpr, r.Value, r.AssertionExpr)
	r.Assertion.acceptAssertionVisitor(pr)
	rr.Render(" ∴ %s❨%s❩ = %s ∵ ", r.AssertionExpr, r.ValueExpr, r.Result)
	r.Rationale.acceptRationaleVisitor(rr)
}

func (rr *rationaleRenderer) VisitIntrinsicRationale(r IntrinsicRationale) {
	pr := &predicateRenderer{
		PredicateExpr: r.PredicateExpr,
		ValueExpr:     r.ValueExpr,
		Form:          affirmativeForm,
		Output:        rr.Output,
	}

	if !r.Result {
		pr.Form = negativeForm
	}

	r.Predicate.acceptPredicateVisitor(pr)
}
