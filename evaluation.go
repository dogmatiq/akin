package akin

import "github.com/dogmatiq/akin/internal/reflectx"

// Evaluation is the result of evaluating a value against a [Predicate].
type Evaluation struct {
	// Predicate is the predicate against which the value was tested.
	Predicate Predicate

	// Value is the value that was evaluated.
	Value any

	// IsSatisfied is true if the value satisfied the predicate.
	IsSatisfied bool

	// Description is a human-readable description of the evaluation result.
	Description string

	// Constituents is a list of evaluations that contributed to this
	// evaluation.
	Constituents []ConstituentEvaluation
}

// ConstituentEvaluation is a specialization for an [Evaluation] that is a
// constituent of some parent evaluation.
type ConstituentEvaluation struct {
	Evaluation
}

func satisfied(p Predicate, v any, format string, args ...any) Evaluation {
	return Evaluation{
		Predicate:   p,
		Value:       v,
		IsSatisfied: true,
		Description: reflectx.Sprintf(format, args...),
	}
}

func violated(p Predicate, v any, format string, args ...any) Evaluation {
	return Evaluation{
		Predicate:   p,
		Value:       v,
		Description: reflectx.Sprintf(format, args...),
	}
}
