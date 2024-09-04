package akin

// A Rationale describes the logical reasoning that justifies an [Evaluation].
//
// Within documentation and strings, ∵ (the because symbol) is used to represent
// a rationale. The latin letter R is used throughout the codebase, which should
// not be confused with 𝑹 (mathematical bold italic capital R) which represents
// a [Predicate] in some circumstances.
type Rationale interface {
	// VisitR calls the method on v associated with the rationale's type.
	VisitR(v RVisitor)
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

	// PVacuous is a [Rationale] based on the fact that 𝑷 has no constituent
	// predicates 𝐐₁, 𝐐₂, … 𝐐ₙ. That is, 𝑛 = 0.
	//
	// Such a [Predicate] does not actually describe any criteria, therefore
	// 𝑷❨𝒙❩ is 𝓾 ([Undefined]) for all 𝒙.
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

// VisitR calls the method on v associated with the rationale's type.
func (r PConst) VisitR(v RVisitor) { v.PConst(r) }
func (r PConst) String() string    { return stringR(r) }
func (s *stringer) PConst(PConst)  { write(s, "𝑷 is constant") }

// VisitR calls the method on v associated with the rationale's type.
func (r PVacuous) VisitR(v RVisitor)  { v.PVacuous(r) }
func (r PVacuous) String() string     { return stringR(r) }
func (s *stringer) PVacuous(PVacuous) { write(s, "𝑷 is vacuous") }

// VisitR calls the method on v associated with the rationale's type.
func (r Px) VisitR(v RVisitor) { v.Px(r) }
func (r Px) String() string    { return stringR(r) }

func (s *stringer) Px(r Px) {
	write(
		s,
		"𝒙 ≔ %s, 𝑷 ≔ %s ∴ 𝑷❨𝒙❩ = %s ∵ %s",
		r.X,
		parens(stringP(r.P, canonical)),
		r.Px,
		r.R,
	)
}

// VisitR calls the method on v associated with the rationale's type.
func (r Qx) VisitR(v RVisitor) { v.Qx(r) }
func (r Qx) String() string    { return stringR(r) }

func (s *stringer) Qx(r Qx) {
	write(
		s,
		"𝐐%s ≔ %s ∴ 𝐐%s❨𝒙) = %s ∵ %s",
		subscript(r.N),
		r.Q,
		subscript(r.N),
		r.Qx,
		r.R,
	)
}

// VisitR calls the method on v associated with the rationale's type.
func (r Ax) VisitR(v RVisitor) { v.Ax(r) }
func (r Ax) String() string    { return stringR(r) }

func (s *stringer) Ax(r Ax) {
	if r.Ax {
		r.A.VisitA(s)
	} else {
		s.Form = !s.Form
		r.A.VisitA(s)
		s.Form = !s.Form
	}
}
