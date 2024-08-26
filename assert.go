package akin

import (
	"fmt"
	"testing"
)

// AssertContains fails the test unless s contains v.
func AssertContains(t *testing.T, s Set, v any) {
	t.Helper()

	m := s.eval(v)

	if !m.IsMember {
		if len(m.Against) == 0 {
			panic(fmt.Sprintf(
				"no reason was provided as to why %s is not a member of %s",
				renderV(valueOf(v)),
				s,
			))
		}

		for _, r := range m.Against {
			t.Logf(
				"expected %s to be a member of %s, but %s",
				renderV(valueOf(v)),
				s,
				r,
			)
		}
		t.Fail()
	}
}

// AssertNotContains fails the test if s contains v.
func AssertNotContains(t *testing.T, s Set, v any) {
	t.Helper()

	m := s.eval(v)

	if m.IsMember {
		if len(m.For) == 0 {
			panic(fmt.Sprintf(
				"no reason was provided as to why %s is a member of %s",
				renderV(valueOf(v)),
				s,
			))
		}

		for _, r := range m.For {
			t.Logf(
				"did not expect %s to be a member of %s, but %s",
				renderV(valueOf(v)),
				s,
				r,
			)
		}
		t.Fail()
	}
}
