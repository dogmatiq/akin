package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestOr(t *testing.T) {
	q := IsEqualTo(1)
	r := IsEqualTo(2)

	p := Or(q, r)

	assertInvariants(t, p)
	assertInvariants(t, Or())
	assertInvariants(t, Or(q))

	assertSatisfied(t, p, 1)
	assertSatisfied(t, p, 2)
	assertViolated(t, p, 3)

	assertEquivalent(t, Or(q, r), Or(r, q))
	assertNotEquivalent(t, Or(q), Or(q, r))

	assertReducesTo(t, Bottom, Or())
	assertReducesTo(t, Bottom, Or(Bottom))

	assertReducesTo(t, Top, Or(Top))
	assertReducesTo(t, Top, Or(Top, q))
	assertReducesTo(t, Top, Or(q, Top))

	assertReducesTo(t, q, Or(q))
	assertReducesTo(t, q, Or(q, q))
	assertIsReduced(t, Or(q, r))

	assertReducesTo(t, q, Or(q, Or(q, q)))
	assertReducesTo(t, Or(q, r), Or(q, Or(r)))
	assertReducesTo(t, Or(q, r), Or(q, Or(q, r)))
}
