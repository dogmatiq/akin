package akin

// TypeEq is a [Predicate] and [Attribute] that holds true when 𝒙 has a
// specific [Type].
//
// Type equality is notated using ∈ (element of) and its inverse ∉ (not an
// element of). For example, given 𝑷 ≔ ❨𝒙 ∈ int❩, then 𝑷❨𝒙❩ = 𝓽 if 𝒙 has a
// [Type] of [int].
type TypeEq struct {
	T Type
}

// VisitP calls the method on v associated with the predicate's type.
func (pa TypeEq) VisitP(v PVisitor) {
	v.TypeEq(pa)
}

// VisitA calls the method on v associated with the attribute's type.
func (pa TypeEq) VisitA(v AVisitor) {
	v.TypeEq(pa)
}

func (pa TypeEq) String() string {
	return stringP(pa, canonical)
}

func (s *stringer) TypeEq(pa TypeEq) {
	write(s, "𝒙 {∈|∉} %s", pa.T)
}

func (e *evaluator) TypeEq(pa TypeEq) {
	sameType := e.X.Type() == pa.T

	e.Px = truth(sameType)
	e.R = Ax{A: pa, Ax: sameType}
}
