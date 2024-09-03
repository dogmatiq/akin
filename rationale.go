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
	PIsConst(PIsConst)
	PIsVacuous(PIsVacuous)
	PofX(PofX)
	QofX(QofX)
	AofX(AofX)
}

type (
	// PIsConst is a [Rationale] based on the fact that some [Predicate] 𝑷
	// produces the same result for any 𝒙.
	PIsConst struct{ P Predicate }

	// PIsVacuous is a [Rationale] based on the fact that 𝑷 has no constituent
	// predicates 𝐐₁, 𝐐₂, … 𝐐ₙ. That is, 𝑛 = 0.
	//
	// Such a [Predicate] does not actually describe any criteria, therefore
	// 𝑷❨𝒙❩ is 𝓾 ([Undefined]) for all 𝒙.
	PIsVacuous struct{ P Predicate }

	// PofX is a [Rationale] based on the evaluation result of 𝑷❨𝒙❩. It is the
	// "top-level" rationale for any call to [Eval].
	PofX struct {
		P Predicate
		X any

		// PX is the result of 𝑷❨𝒙❩.
		PX Truth

		// R is the rationale that justifies 𝑷❨𝒙❩.
		R Rationale
	}

	// QofX is a [Rationale] based on the evaluation result of 𝐐ₙ❨𝒙), where
	// 𝐐ₙ is the 𝑛ᵗʰ constituent of the compound 𝑷.
	QofX struct {
		Q Predicate
		X any

		// N is the 1-based index of the constituent predicate 𝐐ₙ within 𝑷.
		N int

		// QX is the result of 𝐐ₙ❨𝒙).
		QX Truth

		// R is the rationale that justifies 𝐐ₙ❨𝒙).
		R Rationale
	}

	// AofX is a [Rationale] based on whether or not some [Attribute] 𝛂 holds
	// true for 𝒙.
	AofX struct {
		A Attribute

		// AX is the result of 𝛂❨𝒙). That is, it is true if 𝛂 holds for 𝒙.
		AX bool
	}
)

// VisitR calls the method on v associated with the rationale's type.
func (r PIsConst) VisitR(v RVisitor)  { v.PIsConst(r) }
func (r PIsConst) String() string     { return stringR(r) }
func (s *stringer) PIsConst(PIsConst) { s.fmt("𝑷 is constant") }

// VisitR calls the method on v associated with the rationale's type.
func (r PIsVacuous) VisitR(v RVisitor)    { v.PIsVacuous(r) }
func (r PIsVacuous) String() string       { return stringR(r) }
func (s *stringer) PIsVacuous(PIsVacuous) { s.fmt("𝑷 is vacuous") }

// VisitR calls the method on v associated with the rationale's type.
func (r PofX) VisitR(v RVisitor) { v.PofX(r) }
func (r PofX) String() string    { return stringR(r) }

func (s *stringer) PofX(r PofX) {
	s.fmt(
		"𝒙 ≔ %v, 𝑷 ≔ %s ∴ 𝑷❨𝒙❩ = %s ∵ %s", // TODO: use %s and add Value type
		r.X,
		r.P,
		r.PX,
		r.R,
	)
}

// VisitR calls the method on v associated with the rationale's type.
func (r QofX) VisitR(v RVisitor) { v.QofX(r) }
func (r QofX) String() string    { return stringR(r) }

func (s *stringer) QofX(r QofX) {
	s.fmt(
		"𝐐%s ≔ %s ∴ 𝐐%s❨𝒙) = %s ∵ %s",
		subscript(r.N),
		r.Q,
		subscript(r.N),
		r.QX,
		r.R,
	)
}

// VisitR calls the method on v associated with the rationale's type.
func (r AofX) VisitR(v RVisitor) { v.AofX(r) }
func (r AofX) String() string    { return stringR(r) }

func (s *stringer) AofX(r AofX) {
	if r.AX {
		r.A.VisitA(s)
	} else {
		r.A.VisitA((*negatedStringer)(s))
	}
}
