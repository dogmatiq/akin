package akin

import "fmt"

// A Rationale describes the logical reasoning that justifies an [Evaluation].
//
// Within documentation and strings, ∵ (the because symbol) is used to represent
// a rationale. The latin letter R is used throughout the codebase, which should
// not be confused with 𝑹 (mathematical bold italic capital R) which represents
// a [Predicate] in some circumstances.
type Rationale interface {
	visit(RVisitor)
}

// RVisitor is an algorithm with logic specific to each [Rationale] type.
type RVisitor interface {
	PConst(PConst)
	PVacuous(PVacuous)
	Px(Px)
	Qx(Qx)
	Ax(Ax)
}

type (
	// PConst is a [Rationale] based on the fact that some [Predicate] 𝑷
	// produces the same result for any 𝒙.
	PConst struct{ P Predicate }

	// PVacuous is a [Rationale] based on the fact that 𝑷 makes no real
	// assertions.
	//
	// For example, if 𝑷 is a compound [Predicate] with no constituent
	// predicates 𝐐₁, 𝐐₂, … 𝐐ₙ. That is, 𝑛 = 0 then 𝑷 is vacuous.
	//
	// The result of 𝑷❨𝒙❩ is 𝓾 ([Undefined]) for all 𝒙 when 𝑷 is vacuous.
	PVacuous struct{ P Predicate }

	// Px is a [Rationale] based on the evaluation result of 𝑷❨𝒙❩. It is the
	// "top-level" rationale for any call to [Eval].
	Px struct {
		P Predicate
		X Value

		// Px is the result of 𝑷❨𝒙❩.
		Px Truth

		// R is the rationale that justifies 𝑷❨𝒙❩.
		R Rationale
	}

	// Qx is a [Rationale] based on the evaluation result of 𝐐ₙ❨𝒙), where
	// 𝐐ₙ is the 𝑛ᵗʰ constituent of the compound 𝑷.
	Qx struct {
		Q Predicate
		X Value

		// N is the 1-based index of the constituent predicate 𝐐ₙ within 𝑷.
		N int

		// Qx is the result of 𝐐ₙ❨𝒙).
		Qx Truth

		// R is the rationale that justifies 𝐐ₙ❨𝒙).
		R Rationale
	}

	// Ax is a [Rationale] based on whether or not some [Attribute] 𝛂 holds
	// true for 𝒙.
	Ax struct {
		A Attribute

		// Ax is the result of 𝛂❨𝒙). That is, it is true if 𝛂 holds for 𝒙.
		Ax bool
	}
)

func (r PConst) visit(v RVisitor)   { v.PConst(r) }
func (r PVacuous) visit(v RVisitor) { v.PVacuous(r) }
func (r Px) visit(v RVisitor)       { v.Px(r) }
func (r Qx) visit(v RVisitor)       { v.Qx(r) }
func (r Ax) visit(v RVisitor)       { v.Ax(r) }

func (r PConst) String() string {
	return "𝑷 is constant"
}

func (r PVacuous) String() string {
	return "𝑷 is vacuous"
}

func (r Px) String() string {
	return fmt.Sprintf(
		"%s ≔ %s, 𝑷 ≔ %s ∴ 𝑷❨%s❩ = %s ∵ %s",
		r.X.Expr(),
		r.X,
		parens(stringP(r.P, affirmative)),
		r.X.Expr(),
		r.Px,
		r.R,
	)
}

func (r Qx) String() string {
	return fmt.Sprintf(
		"𝐐%s ≔ %s ∴ 𝐐%s❨𝒙) = %s ∵ %s",
		subscript(r.N),
		r.Q,
		subscript(r.N),
		r.Qx,
		r.R,
	)
}

func (r Ax) String() string {
	return stringA(r.A, form(r.Ax))
}
