package akin

import (
	"fmt"
)

// ValueEquivalence is a [Property] that is satisfied when 𝑥 is equivalent to
// some value.
//
// "Value" in this context refers to the general concept of a value. For
// example, the the value "1" means the number one, as opposed to the result of
// the int(1) expression in Go.
type ValueEquivalence struct {
	Description string
}

// Format implements the [fmt.Formatter] interface.
func (p ValueEquivalence) Format(s fmt.State, v rune) {
	format(p, s, v)
}
func (p ValueEquivalence) hide() any {
	type T = ValueEquivalence
	type ValueEquivalence T
	return ValueEquivalence(p)
}

func (p ValueEquivalence) formal() string {
	return sprintf("𝑥 = %s", p.Description)
}

func (p ValueEquivalence) human() string {
	return sprintf("𝑥 is %s", p.Description)
}

func (p ValueEquivalence) inverse() string {
	return sprintf("𝑥 is not %s", p.Description)
}

func (p ValueEquivalence) visitProperty(v PropertyVisitor) {
	v.VisitValueEquivalenceProperty(p)
}
