package akin_test

import (
	"testing"

	. "github.com/dogmatiq/akin"
	. "github.com/dogmatiq/akin/internal/testx"
)

func TestIdentity(t *testing.T) {
	for _, c1 := range AllCases {
		t.Run(c1.Name, func(t *testing.T) {
			p := Is(c1.X)

			AssertTrue(t, p, c1.X)

			for _, c2 := range AllCases {
				if c1 != c2 {
					t.Run(c2.Name, func(t *testing.T) {
						AssertFalse(t, p, c2.X)
					})
				}
			}
		})
	}

	t.Run("slices", func(t *testing.T) {
		x := []int{1, 2, 3}
		p := Is(x)

		AssertTrue(t, p, x)
		AssertTrue(t, p, x[0:])
		AssertFalse(t, p, x[:0])
		AssertFalse(t, p, x[:2])
		AssertFalse(t, p, x[1:])
	})

	t.Run("maps", func(t *testing.T) {
		x := map[int]int{}
		y := map[int]int{}
		var z map[int]int

		p := Is(x)

		AssertTrue(t, p, x)
		AssertFalse(t, p, y)
		AssertFalse(t, p, z)
	})

	t.Run("arrays", func(t *testing.T) {
		f1 := func() { panic("has a side-effect") }
		f2 := func() { /* has no side effect */ }

		x := [1]func(){f1}
		y := [1]func(){f1}
		z := [1]func(){f2}

		p := Is(x)

		AssertTrue(t, p, x)
		AssertTrue(t, p, y)
		AssertFalse(t, p, z)
	})

	t.Run("structs", func(t *testing.T) {
		f1 := func() { panic("has a side-effect") }
		f2 := func() { /* has no side effect */ }

		x := struct{ v func() }{f1}
		y := struct{ v func() }{f1}
		z := struct{ v func() }{f2}

		p := Is(x)

		AssertTrue(t, p, x)
		AssertTrue(t, p, y)
		AssertFalse(t, p, z)
	})

	t.Run("functions", func(t *testing.T) {
		x := func() { panic("has a side-effect") }
		y := func() { /* has no side effect */ }
		var z func()

		p := Is(x)

		AssertTrue(t, p, x)
		AssertFalse(t, p, y)
		AssertFalse(t, p, z)
	})
}
