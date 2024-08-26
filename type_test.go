package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestType(t *testing.T) {
	set := HasType[int]()

	AssertContains(t, set, 0)
	AssertNotContains(t, set, uint(0))
	AssertNotContains(t, set, float64(0))
	AssertNotContains(t, set, "0")
}
