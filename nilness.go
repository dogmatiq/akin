package akin

// Nilness is a [Predicate] that is satisfied when 𝒙 has "nilness" equal to the
// predicate's value.
//
// There are two possible values; [IsNil] and [IsNonNil].
type Nilness bool

const (
	// IsNil is a [Predicate] that evaluates to [True] when 𝒙 is nil,
	// regardless of its type.
	IsNil Nilness = true

	// IsNonNil is a [Predicate] that evaluates to [False] when 𝒙 is nil,
	// regardless of its type.
	IsNonNil Nilness = false
)

// VisitP calls the method on v associated with the predicate's type.
func (p Nilness) VisitP(v PVisitor) {
	v.Nilness(p)
}

func (p Nilness) String() string {
	return stringP(p)
}

func (s *identity) Nilness(p Nilness) {
	if p {
		*s = "𝒙 ≍ nil"
	} else {
		*s = "𝒙 ≭ nil"
	}
}

func (e *evaluator) Nilness(p Nilness) {
	wantNil := bool(p)
	gotNil := e.X.isNil()
	e.Px = truth(gotNil == wantNil)

	if e.X.Type().isNilable() {
		e.R = Ax{
			A:  ValueEq{"nil"},
			Ax: gotNil,
		}
	} else {
		e.R = Ax{
			A:  TypeEq{e.X.Type()},
			Ax: true,
		}
	}
}
