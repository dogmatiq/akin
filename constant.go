package akin

const (
	// Top is a [Predicate] that is satisfied by any value.
	Top constant = true

	// Bottom is a [Predicate] that is not satisfied by any values.
	Bottom constant = false
)

type constant bool

var (
	_ Predicate = Top
	_ Predicate = Bottom
)

func (p constant) String() string {
	if p {
		return "⊤"
	}
	return "⊥"
}

func (p constant) Eval(v any) Evaluation {
	if p {
		return satisfied(p, v, "all values satisfy %s", p)
	}
	return violated(p, v, "no values satisfy %s", p)
}

func (p constant) Is(q Predicate) bool {
	return p == q
}

func (p constant) Reduce() Predicate {
	return p
}
