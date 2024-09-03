package akin

// Nilness is a [Predicate] that is satisfied when 𝒙 has "nilness" equal to the
// predicate's value. See [IsNil] and [IsNonNil].
type Nilness bool

const (
	// IsNil is a [Predicate] that is satisfied when 𝒙 is nil, regardless of
	// its type.
	IsNil Nilness = true

	// IsNonNil is a [Predicate] that is satisfied when 𝒙 is non-nil,
	// regardless of its type.
	IsNonNil Nilness = false
)

func (p Nilness) visitP(v PVisitor) { v.Nilness(p) }
func (p Nilness) String() string    { return "𝒙 is nil" }
func (p Nilness) NString() string   { return "𝒙 is non-nil" } //revive:disable-line:exported

func (e *evaluator) Nilness(p Nilness) {
	wantNil := bool(p)
	gotNil := e.Value.isNil()
	e.Result = tern(gotNil == wantNil)

	if e.Value.Type().isNilable() {
		e.Rationale = PropertyRationale{
			Property: ValueEquivalence{"nil"},
			Holds:    gotNil,
		}
	} else {
		e.Rationale = PropertyRationale{
			Property: TypeEquivalence{e.Value.Type()},
			Holds:    true,
		}
	}
}
