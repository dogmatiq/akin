package akin

// Is returns a [Predicate] that holds true when 𝒙 has a specific [Type].
func Is[T any]() TypeEq {
	return TypeEq{T: typeFor[T]()}
}

// TypeEq is a [Predicate] and [Attribute] that holds true when 𝒙 has a
// specific [Type].
//
// Type equality is notated using ∈ (element of) and its inverse ∉ (not an
// element of). For example, given 𝑷 ≔ ❨𝒙 ∈ int❩, then 𝑷❨𝒙❩ = 𝓽 if 𝒙 has a
// [Type] of [int].
type TypeEq struct {
	T Type
}

func (pa TypeEq) visitP(v PVisitor)  { v.TypeEq(pa) }
func (pa TypeEq) visitA(v AVisitor)  { v.TypeEq(pa) }
func (pa TypeEq) String() string     { return stringP(pa, affirmative) }
func (s *stringer) TypeEq(pa TypeEq) { render(s, "𝒙 {∈|∉} %s", pa.T) }

func (e *evaluator) TypeEq(pa TypeEq) {
	t := e.X.Type()

	if t == pa.T {
		e.Px = True
		e.R = Ax{A: pa, Ax: true}
	} else {
		e.Px = False
		e.R = Ax{A: TypeEq{t}, Ax: true}
	}
}
