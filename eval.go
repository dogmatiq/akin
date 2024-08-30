package akin

import (
	"fmt"
	"reflect"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// Evaluation is the result of evaluating a value 𝑥 against the [Predicate] 𝑷.
type Evaluation struct {
	P           Predicate
	X           any
	IsSatisfied bool
	Reason      Reason
}

// Format implements the [fmt.Formatter] interface.
func (e Evaluation) Format(s fmt.State, v rune) {
	format(e, s, v)
}

func (e Evaluation) hide() any {
	type T = Evaluation
	type Evaluation T
	return T(e)
}

func (e Evaluation) formal() string {
	return sprintf(
		"𝑷❨𝑥❩ ≔ %s; 𝑥 ≔ %s ∴ 𝑷❨𝑥❩ = %t ∵ %+s",
		e.P,
		reflectx.ValueOf(e.X),
		e.IsSatisfied,
		e.Reason,
	)
}

func (e Evaluation) human() string {
	return sprintf(
		"%s %s %s because %+s",
		reflectx.ValueOf(e.X),
		choose(e.IsSatisfied, "satisfies", "violates"),
		e.P,
		e.Reason,
	)
}

func eval(p Predicate, x reflect.Value) Evaluation {
	e := &evaluator{
		X: x,
		T: x.Type(),

		Evaluation: Evaluation{
			P: p,
			X: x.Interface(),
		},
	}

	VisitPredicate(p, e)

	if e.Reason == nil {
		panic("no reason provided")
	}

	return e.Evaluation
}

type evaluator struct {
	Evaluation

	X reflect.Value
	T reflect.Type
}

func (e *evaluator) SetReason(satisfied bool, r Reason) {
	e.IsSatisfied = satisfied
	e.Reason = r
}

func (e *evaluator) SetProperty(satisfied bool, p Property) {
	e.IsSatisfied = satisfied

	if satisfied {
		e.Reason = PropertySatisfied{p}
	} else {
		e.Reason = PropertyViolated{p}
	}
}
