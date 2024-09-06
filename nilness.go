package akin

// Nilness is a [PredicateP] that is satisfied when 𝒙 has "nilness" equal to the
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

func (p Nilness) acceptPredicateVisitor(v PredicateVisitor) { v.VisitNilness(p) }
func (p Nilness) acceptAssertionVisitor(v AssertionVisitor) { v.VisitNilness(p) }
func (p Nilness) String() string                            { return toDefaultString(p) }

func (r *renderer) VisitNilness(p Nilness) {
	if p {
		r.render("𝒙 {=|≠} nil")
	} else {
		r.render("𝒙 {≠|=} nil")
	}
}

func (e *evaluator) VisitNilness(p Nilness) {
	wantNil := bool(p)
	gotNil := e.Value.isNil()
	e.Result = asResult(gotNil == wantNil)

	if e.Value.Type().isNilable() {
		e.Rationale = IntrinsicRationale{
			Predicate: ValueEq{"nil"},
			Value:     e.Value,
			Result:    gotNil,
		}
	} else {
		e.Rationale = IntrinsicRationale{
			Predicate: TypeEq{e.Value.Type()},
			Value:     e.Value,
			Result:    true,
		}
	}
}
