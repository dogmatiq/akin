package akin

// ValueEq is an [Attribute] that holds when 𝒙 is equal to some specific
// abstract value.
//
// Abstract in this context means that the value is conceptual, such as the
// number one, versus a specific Go value, such as float64(1.0).
//
// Value equality is notated using ≍ (equivalent to) and its inverse ≭ (not
// equivalent to). For example, given 𝛂 ≔ ❨𝒙 ≍ 1❩, then 𝛂❨𝒙❩ = 𝓽 if 𝒙 is
// the number one.
type ValueEq struct {
	// Repr is a human-readable representation of the value.
	Repr string
}

func (a ValueEq) visit(v AVisitor)    { v.ValueEq(a) }
func (a ValueEq) String() string      { return stringA(a, affirmative) }
func (s *stringer) ValueEq(a ValueEq) { render(s, "𝒙 {≍|≭} %s", a.Repr) }
