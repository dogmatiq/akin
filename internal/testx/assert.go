package testx

import (
	"fmt"
	"testing"

	"github.com/dogmatiq/akin"
)

// AssertRationale asserts that 𝑷❨𝒙❩ produces a specific rationale.
func AssertRationale(t *testing.T, p akin.Assertion, x any, want string) {
	t.Helper()

	_, r := akin.Eval(p, x)
	got := fmt.Sprint(r)

	if got == want {
		t.Logf("\x1b[32m✔\x1b[0m %s", want)
	} else {
		t.Logf("  \x1b[2;31m- \x1b[0m%s", got)
		t.Errorf("\x1b[0;31m✘ \x1b[2;32m+ \x1b[0m%s", want)
	}
}

// AssertTrue asserts that 𝑷❨𝒙❩ evaluates to 𝓽.
func AssertTrue(t *testing.T, p akin.Assertion, x any) {
	t.Helper()
	assert(t, akin.True, p, x)
}

// AssertFalse asserts that 𝑷❨𝒙❩ evaluates to 𝓯.
func AssertFalse(t *testing.T, p akin.Assertion, x any) {
	t.Helper()
	assert(t, akin.False, p, x)
}

func assert(
	t *testing.T,
	expect akin.Result,
	p akin.Assertion,
	x any,
) {
	t.Helper()

	result, rat := akin.Eval(p, x)

	if result == expect {
		t.Logf("\x1b[32m✔\x1b[0m %s", rat)
	} else {
		t.Errorf("\x1b[31m✘\x1b[0m %s", rat)
	}
}
