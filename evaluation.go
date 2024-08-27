package akin

import "fmt"

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
		Reason:      fmt.Sprintf(format, args...),
	}
}

func violated(p Predicate, v any, format string, args ...any) Evaluation {
	return Evaluation{
		Predicate: p,
		Value:     v,
		Reason:    fmt.Sprintf(format, args...),
	}
}
