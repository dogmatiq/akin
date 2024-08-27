package akin

const (
	// Universe is the [Set] of all possible values.
	Universe extremum = true

	// Empty is a [Set] containing no values.
	Empty extremum = false
)

// An extremum is a [Set] that contains either all values or no values.
type extremum bool

var (
	_ Set = Universe
	_ Set = Empty
)

func (s extremum) String() string {
	if s {
		return "Ω"
	}
	return "∅"
}

func (s extremum) Contains(any) bool {
	return bool(s)
}

func (s extremum) eval(v any) evaluation {
	return newEvaluation(
		s,
		v,
		bool(s),
		constant(s),
	)
}

// constant is a [predicate] that always returns the same result.
type constant bool

func (p constant) String(bool) string {
	if p {
		return "everything is a member of Ω"
	}
	return "nothing is a member of ∅"
}
