package akin

// Const is an [Assertion] that produces the same result regardless of the
// [Value].
//
// There are two possible values; [Top] and [Bottom], denoted ⊤ and ⊥,
// respectively.
//
// - [Top] is [True] for any 𝒙, that is ⊤❨𝒙) ≔ 𝓽 - [Bottom] is [False] for
// any 𝒙, that is ⊥❨𝒙) ≔ 𝓯
type Const bool

const (
	// Top (denoted ⊤) is a [Predicate] that evaluates to [True] for any 𝒙.
	Top Const = true

	// Bottom (denoted ⊥) is a [Predicate] that evaluates to [False] for any 𝒙.
	Bottom Const = false
)

func (p Const) acceptPredicateVisitor(v PredicateVisitor) { v.VisitConst(p) }
func (p Const) acceptAssertionVisitor(v AssertionVisitor) { v.VisitConst(p) }

func (p Const) String() string {
	if p {
		return "⊤"
	}
	return "⊥"
}

func (s *renderer) VisitConst(p Const) {
	if s.Form == negative {
		p = !p
	}
	s.Output = p.String()
}

func (e *evaluator) VisitConst(p Const) {
	e.Result = asResult(p)
	e.Rationale = ConstRationale{p}
}
