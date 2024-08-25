package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func assertAkin(t *testing.T, spec, v any) {
	t.Helper()
	if Test(spec, v) != nil {
		t.Errorf("expected %T(%t) to be akin to %T(%t)", spec, spec, v, v)
	}
}

func assertNotAkin(t *testing.T, spec, v any) {
	t.Helper()
	if Test(spec, v) == nil {
		t.Errorf("did not expect %T(%t) to be akin to %T(%t)", spec, spec, v, v)
	}
}
