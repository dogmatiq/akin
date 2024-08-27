package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestExtremum(t *testing.T) {
	for _, c := range all {
		t.Run(c.Name, func(t *testing.T) {
			AssertIsMember(t, Universe, c.Value)
			AssertIsNotMember(t, Empty, c.Value)
		})
	}
}
