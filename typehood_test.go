package akin_test

import (
	"io"
	"testing"

	. "github.com/dogmatiq/akin"
	. "github.com/dogmatiq/akin/internal/testx"
)

func TestTypehood(t *testing.T) {
	t.Run("concrete type", func(t *testing.T) {
		p := Is[int]()

		AssertIsReduced(t, p)

		AssertTrue(t, p, 0)
		AssertFalse(t, p, uint(0))
		AssertFalse(t, p, float64(0))
		AssertFalse(t, p, "0")
	})

	t.Run("interface type", func(t *testing.T) {
		p := Is[error]()

		AssertIsReduced(t, p)

		AssertTrue(t, p, io.EOF)
		AssertFalse(t, p, 1)
		AssertFalse(t, p, nil)
	})
}
