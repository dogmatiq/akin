package akin

import (
	"slices"
	"testing"
)

// AssertContains fails the test unless s contains v.
func AssertContains(t *testing.T, s Set, v any, reasons ...string) {
	t.Helper()

	m := s.Eval(v)
	if !m.IsMember {
		t.Fatalf("expected %T(%#v) to be a member of %s: %s", v, v, s, m.Reason)
	}

	if len(reasons) != 0 && !slices.Contains(reasons, m.Reason) {
		t.Fatalf("unexpected reason: %s", m.Reason)
	}
}

// AssertNotContains fails the test if s contains v.
func AssertNotContains(t *testing.T, s Set, v any, reasons ...string) {
	t.Helper()

	m := s.Eval(v)

	if m.IsMember {
		t.Fatalf("expected %T(%#v) to not be a member of %s: %s", v, v, s, m.Reason)
	}

	if len(reasons) != 0 && !slices.Contains(reasons, m.Reason) {
		t.Fatalf("unexpected reason: %s", m.Reason)
	}
}
