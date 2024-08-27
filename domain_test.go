package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestDomain(t *testing.T) {
	set := DomainFor[int]()

	AssertIsMember(t, set, 0)
	AssertIsNotMember(t, set, uint(0))
	AssertIsNotMember(t, set, float64(0))
	AssertIsNotMember(t, set, "0")
}
