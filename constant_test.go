package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestConstant(t *testing.T) {
	for _, c := range all {
		t.Run(c.Name, func(t *testing.T) {
			AssertSatisfied(t, Anything, c.Value)
			AssertViolated(t, Nothing, c.Value)
		})
	}
}
