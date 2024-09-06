package akin

import "fmt"

// Eval evaluates 𝑷❨𝒙❩, returning the [Result] and a [Rationale] for that result.
func Eval(p Assertion, x any) (Result, Rationale) {
	e := evaluator{
		Predicate: p,
		Value:     valueOf(x),
		ValueExpr: varX,
	}

	p.acceptAssertionVisitor(&e)

	if e.Rationale == nil {
		panic("no rationale provided")
	}

	return e.Result, AssertionRationale(e)
}

type evaluator AssertionRationale

// Result is a three-valued logic type.
//
// It has three possible values; [True], [False] and [Undefined], denoted using
// the mathematical bold script small letters 𝓽, 𝓯 and 𝓾, respectively.
type Result struct{ v int }

var (
	// Undefined (or 𝓾) is a [Result] that is neither [True] nor [False].
	Undefined = Result{}

	// True (or 𝓽) is the [Result] that is equivalent to the boolean true.
	True = Result{1}

	// False (or 𝓯) is the [Result] that is equivalent to the boolean false.
	False = Result{-1}
)

func asResult[T ~bool](v T) Result {
	if v {
		return True
	}
	return False
}

func (t Result) String() string {
	switch t {
	case True:
		return "𝓽"
	case False:
		return "𝓯"
	default:
		return "𝓾"
	}
}

// A Rationale describes the logical reasoning that justifies an [Evaluation].
//
// The ∵ symbol (because) is used to represent a rationale.
type Rationale interface {
	fmt.Stringer

	acceptRationaleVisitor(RationaleVisitor)
}

// RationaleVisitor is an algorithm with logic specific to each [Rationale] type.
type RationaleVisitor interface {
	VisitConstRationale(ConstRationale)
	VisitVacuousRationale(VacuousRationale)
	VisitAssertionRationale(AssertionRationale)
	VisitAttributeRationale(IntrinsicRationale)
}

type (
	// ConstRationale is the [Rationale] provided when a [Predicate] always
	// produces the same [Result] regardless of the [Value].
	ConstRationale struct {
		Predicate Predicate
	}

	// VacuousRationale is the [Rationale] provided when a [Predicate] makes no
	// meaningful assertion about the [Value].
	//
	// For example, a compound predicate with no constituent predicates is
	// considered vacuous. The [Result] of a vacuous predicate is always
	// [Undefined].
	VacuousRationale struct {
		Predicate Predicate
	}

	// AssertionRationale is the [Rationale] provided when a [Result] is
	// determined by evaluation of an [Assertion].
	AssertionRationale struct {
		Predicate Assertion
		Value     Value
		ValueExpr ValueExpr
		Result    Result
		Rationale Rationale
	}

	// IntrinsicRationale is the [Rationale] provided when a [Result] is derived
	// from a [Predicate] that is not an [Assertion].
	//
	// Such predicates describe intrinsic attributes of a value which may be
	// abstract in nature.
	IntrinsicRationale struct {
		Predicate Predicate
		Value     Value
		ValueExpr ValueExpr
		Result    bool
	}
)

func (r ConstRationale) acceptRationaleVisitor(v RationaleVisitor)     { v.VisitConstRationale(r) }
func (r VacuousRationale) acceptRationaleVisitor(v RationaleVisitor)   { v.VisitVacuousRationale(r) }
func (r AssertionRationale) acceptRationaleVisitor(v RationaleVisitor) { v.VisitAssertionRationale(r) }
func (r IntrinsicRationale) acceptRationaleVisitor(v RationaleVisitor) { v.VisitAttributeRationale(r) }

func (r ConstRationale) String() string {
	return "𝑷 is constant"
}

func (r VacuousRationale) String() string {
	return "𝑷 is vacuous"
}

func (r AssertionRationale) String() string {
	p := toString(r.Predicate, r.ValueExpr, affirmative)

	return fmt.Sprintf(
		"%s ≔ %s, 𝑷 ≔ %s ∴ 𝑷❨%s❩ = %s ∵ %s",
		r.ValueExpr,
		r.Value,
		parens(p),
		r.ValueExpr,
		r.Result,
		r.Rationale,
	)
}

func (r IntrinsicRationale) String() string {
	f := affirmative
	if !r.Result {
		f = negative
	}

	return toString(r.Predicate, r.ValueExpr, f)
}
