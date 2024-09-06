package akin

// TypeEq is an [Attribute] that holds when 𝒙 has a specific [Type].
//
// Type equality is notated using ⦂ (z notation type colon). For example, given
// 𝑷 ≔ ❨𝒙 ⦂ int❩, then 𝑷❨𝒙❩ = 𝓽 if 𝒙 has a [Type] of [int].
type TypeEq struct {
	T Type
}

func (p TypeEq) acceptPredicateVisitor(v PredicateVisitor) { v.VisitTypeEq(p) }
func (p TypeEq) String() string                            { return predicateToString(p) }

func (r *predicateRenderer) VisitTypeEq(p TypeEq) {
	r.Render("𝒙 ⦂ %s", p.T)
}
