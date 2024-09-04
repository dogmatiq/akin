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

func (p Nilness) visit(v PVisitor)    { v.Nilness(p) }
func (p Nilness) String() string      { return stringP(p, affirmative) }
func (s *stringer) Nilness(p Nilness) { renderNegatable(s, p, "𝒙 {≍|≭} nil") }

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
