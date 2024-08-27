package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestTo_nil(t *testing.T) {
	set := To(nil)

	for _, c := range nils {
		t.Run(c.Name, func(t *testing.T) {
			AssertIsMember(t, set, c.Value)
		})
	}

	for _, c := range nonNils {
		t.Run(c.Name, func(t *testing.T) {
			AssertIsNotMember(t, set, c.Value)
		})
	}
}
