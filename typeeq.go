package akin

// TypeEq is an [Attribute] that holds when 𝒙 has a specific [Type].
//
// Type equality is notated using ⦂ (z notation type colon). For example, given
// 𝑷 ≔ ❨𝒙 ⦂ int❩, then 𝑷❨𝒙❩ = 𝓽 if 𝒙 has a [Type] of [int].
type TypeEq struct {
	T Type
}

func (a TypeEq) acceptPredicateVisitor(v PredicateVisitor) { v.VisitTypeEq(a) }
func (a TypeEq) String() string                            { return toDefaultString(a) }

func (r *renderer) VisitTypeEq(a TypeEq) {
	r.render("𝒙 ⦂ %s", a.T)
}
