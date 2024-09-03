package akin

// Const is a [Predicate] that produces the same result for all 𝒙.
//
// There are two possible values; [Top] and [Bottom], denoted ⊤ and ⊥,
// respectively.
//
// - [Top] is [True] for any 𝒙, that is ⊤❨𝒙) ≔ 𝓽
// - [Bottom] is [False] for any 𝒙, that is ⊥❨𝒙) ≔ 𝓯
type Const bool

const (
	// Top (denoted ⊤) is a [Predicate] that is [Satisfied] for any 𝒙.
	Top Const = true

	// Bottom (denoted ⊥) is a [Predicate] that is [Violated] for any 𝒙.
	Bottom Const = false
)

// VisitP calls the method on v associated with the predicate's type.
func (p Const) VisitP(v PVisitor) { v.Const(p) }
func (p Const) String() string    { return stringP(p) }

func (s *stringer) Const(p Const) {
	if p {
		*s = "⊤"
	} else {
		*s = "⊥"
	}
}

func (e *evaluator) Const(p Const) {
	e.PX = truth(p)
	e.R = PIsConst{p}
}
