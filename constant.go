package akin

const (
	// Anything is a [Predicate] that is satisfied by any value.
	Anything constant = true

	// Nothing is a [Predicate] that is not satisfied by any values.
	Nothing constant = false
)

type constant bool

var (
	_ Predicate = Anything
	_ Predicate = Nothing
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
	if q, ok := q.(constant); ok {
		return p == q
	}
	return false
}

func (p constant) Simplify() (Predicate, bool) {
	return p, false
}
