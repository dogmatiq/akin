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

func (s constant) String() string {
	if s {
		return "⊤"
	}
	return "⊥"
}

func (s constant) Eval(v any) Evaluation {
	if s {
		return satisfied(s, v, "all values satisfy %s", s)
	}
	return violated(s, v, "no values satisfy %s", s)
}
