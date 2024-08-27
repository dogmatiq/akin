package akin

import (
	"testing"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// AssertIsMember fails the test unless v is a member of s.
func AssertIsMember(t *testing.T, s Set, v any) {
	t.Helper()

	t.Logf(
		"asserting that %q is a member of %q",
		renderTV(reflectx.ValueOf(v)),
		s,
	)

	e := s.eval(v)
	if !e.IsMember {
		t.Fail()
	}

	logEvaluation(t, e)

}

// AssertIsNotMember fails the test if v is a member of s.
func AssertIsNotMember(t *testing.T, s Set, v any) {
	t.Helper()

	t.Logf(
		"asserting that %q is not a member of %q",
		renderTV(reflectx.ValueOf(v)),
		s,
	)

	e := s.eval(v)
	if e.IsMember {
		t.Fail()
	}

	logEvaluation(t, e)
}

func logEvaluation(t *testing.T, e evaluation) {
	t.Helper()

	summarize := true

	for _, p := range e.Predicates {
		if p.Set == e.Set {
			summarize = false
		}

		if p.Satisfied {
			t.Logf(
				" - member of %s because %s",
				p.Set,
				p.Predicate.String(false),
			)
		} else {
			t.Logf(
				" - not a member of %s because %s",
				p.Set,
				p.Predicate.String(true),
			)
		}
	}

	if summarize {
		if e.IsMember {
			t.Logf(" - and therefore it is a member of %s", e.Set)
		} else {
			t.Logf(" - and therefore it is not a member of %s", e.Set)
		}
	}
}
