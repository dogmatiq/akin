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
func (p Nilness) String() string                            { return predicateToString(p) }

func (pr *predicateRenderer) VisitNilness(p Nilness) {
	if p {
		pr.Render("𝒙 {=|≠} nil")
	} else {
		pr.Render("𝒙 {≠|=} nil")
	}
}

func (e *evaluator) VisitNilness(p Nilness) {
	wantNil := bool(p)
	gotNil := e.Value.isNil()
	e.Result = asResult(gotNil == wantNil)

	if e.Value.Type().isNilable() {
		e.Rationale = IntrinsicRationale{
			ValueEq{"nil"},
			PrimeExpr{e.PredicateExpr, 0},
			e.Value,
			e.ValueExpr,
			gotNil,
		}
	} else {
		e.Rationale = IntrinsicRationale{
			TypeEq{e.Value.Type()},
			PrimeExpr{e.PredicateExpr, 0},
			e.Value,
			e.ValueExpr,
			true,
		}
	}
}
