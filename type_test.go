package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestOfType(t *testing.T) {
	p := OfType[int]()

	AssertSatisfied(t, p, 0)
	AssertViolated(t, p, uint(0))
	AssertViolated(t, p, float64(0))
	AssertViolated(t, p, "0")
}
