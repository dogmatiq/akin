package akin

import (
	"testing"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// AssertSatisfied fails the test unless v satisfies p.
func AssertSatisfied(t *testing.T, p Predicate, v any) {
	t.Helper()

	e := p.Eval(v)

	if e.IsSatisfied {
		t.Logf(
			"as expected, %q satisfies %q, because %s",
			render(reflectx.ValueOf(v)),
			p,
			e.Reason,
		)
	} else {
		t.Errorf(
			"expected %q to satisfy %q, but %s",
			render(reflectx.ValueOf(v)),
			p,
			e.Reason,
		)
	}

	if e.Reason == "" {
		panic("no reason provided")
	}
}

// AssertViolated fails the test if v satisfies p.
func AssertViolated(t *testing.T, p Predicate, v any) {
	t.Helper()

	e := p.Eval(v)

	if e.IsSatisfied {
		t.Errorf(
			"expected %q to violate %q, but %s",
			render(reflectx.ValueOf(v)),
			p,
			e.Reason,
		)
	} else {
		t.Logf(
			"as expected, %q violates %q because %s",
			render(reflectx.ValueOf(v)),
			p,
			e.Reason,
		)
	}

	if e.Reason == "" {
		panic("no reason provided")
	}
}
