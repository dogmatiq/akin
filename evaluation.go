package akin

import "github.com/dogmatiq/akin/internal/reflectx"

// Evaluation is the result of evaluating a value against a [Predicate].
type Evaluation struct {
	Predicate    Predicate
	Value        any
	IsSatisfied  bool
	Reason       string
	Constituents []Evaluation
}

func satisfied(p Predicate, v any, format string, args ...any) Evaluation {
	return Evaluation{
		Predicate:   p,
		Value:       v,
		IsSatisfied: true,
		Reason:      reflectx.Sprintf(format, args...),
	}
}

func violated(p Predicate, v any, format string, args ...any) Evaluation {
	return Evaluation{
		Predicate: p,
		Value:     v,
		Reason:    reflectx.Sprintf(format, args...),
	}
}
