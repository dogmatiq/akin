package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestOfType(t *testing.T) {
	p := OfType[int]()

	assertSatisfied(t, p, 0)
	assertViolated(t, p, uint(0))
	assertViolated(t, p, float64(0))
	assertViolated(t, p, "0")
}
