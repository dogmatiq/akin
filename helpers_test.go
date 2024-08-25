package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func assertAkin(t *testing.T, spec, v any) {
	t.Helper()
	if err := Test(spec, v); err != nil {
		t.Errorf("expected %T(%#v) to be akin to %T(%#v): %s", spec, spec, v, v, err)
	}
}

func assertNotAkin(t *testing.T, spec, v any) {
	t.Helper()
	if Test(spec, v) == nil {
		t.Errorf("did not expect %T(%#v) to be akin to %T(%#v)", spec, spec, v, v)
	}
}
