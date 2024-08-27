package akin

import "github.com/dogmatiq/akin/internal/reflectx"

const (
	// Nil is the [Set] of all nil values.
	Nil nilness = true

	// NonNil is the [Set] of all non-nil values.
	NonNil nilness = false
)

type nilness bool

var (
	_ Set = Nil
	_ Set = NonNil
)

func (s nilness) String() string {
	op := "=="
	if !s {
		op = "!="
	}

	return "{ x | x " + op + " nil }"
}

func (s nilness) Contains(v any) bool {
	r := reflectx.ValueOf(v)
	return bool(s) == reflectx.IsNil(r)
}

func (s nilness) eval(v any) evaluation {
	return newEvaluation(
		s,
		v,
		s.Contains(v),
		isNil{},
	)
}

type isNil struct{}

func (r isNil) String(inverse bool) string {
	if inverse {
		return "it is non-nil"
	}
	return "it is nil"
}
