package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestInhabitance(t *testing.T) {
	p := InhabitsType[int]()

	assertSatisfied(t, p, 0)
	assertViolated(t, p, uint(0))
	assertViolated(t, p, float64(0))
	assertViolated(t, p, "0")
}
