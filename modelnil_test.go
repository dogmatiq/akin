package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
)

func TestTo_nil(t *testing.T) {
	set := To(nil)

	for _, v := range nils {
		AssertContains(t, set, v)
	}

	for _, v := range nonNils {
		AssertNotContains(t, set, v)
	}
}
