package akin

import (
	"github.com/dogmatiq/akin/internal/fmtx"
	"github.com/dogmatiq/akin/internal/reflectx"
)

// Evaluation is the result of evaluating a [Predicate] against a [Value].
type Evaluation struct {
	name      string
	Predicate Predicate
	Value     Value
	Result    Tern
	Rationale Rationale
}

func (e Evaluation) String() string {
	return fmtx.F(
		"%s ≔ %s, 𝒙 ≔ %s ∴ %s❨𝒙❩ = %s ∵ %s",
		e.name,
		e.Predicate,
		e.Value,
		e.name,
		e.Result,
		e.Rationale,
	)
}

// Eval returns the evaluating a [Predicate] against a value.
func Eval(p Predicate, x any) Evaluation {
	v := Value{
		reflectx.ValueOf(x),
	}

	var n namer
	e := Evaluation{
		name:      n.Next(),
		Predicate: p,
		Value:     v,
	}

	VisitP(
		p,
		&evaluator{
			Name:       n.Next,
			Evaluation: &e,
		},
	)

	if e.Rationale == nil {
		panic("no rationale provided")
	}

	return e
}

type namer rune

func (n *namer) Next() string {
	*n++
	return string('𝑷' + *n - 1)
}

type evaluator struct {
	*Evaluation
	Name func() string
}
