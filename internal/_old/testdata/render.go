package akin

import (
	"fmt"
	"io"
)

type renderer struct {
}

// renderContext describes the context in which a [renderable] is rendered.
type renderContext struct {
	io.Writer

	// Inverse indicates that rendered output should represent the inverse.
	Inverse bool

	// Verbatim indicates that the rendered output should not be simplified in
	// any way.
	Verbatim bool

	// Verbose indicates that the output should be in a detailed human-readable
	// form.
	Verbose bool
}

var errNoInverse = fmt.Errorf("cannot render inverse")

type renderable interface {
	// render(renderContext) error

	// hide returns the value represented as a type with the same underlying
	// type as the original value, but with no methods. This allows us to format
	// the value without recursion into the [fmt.Stringer.String] or
	// [fmt.Formatter.Format] methods.
	//
	// See https://github.com/golang/go/issues/51195#issuecomment-1563538796
	hide() any
}

func format(r renderable, s fmt.State, v rune) {
	if v == 's' {
		if err := r.render(renderContext{
			Writer:  s,
			Verbose: s.Flag('+'),
		}); err != nil {
			panic(err)
		}
		return
	}

	f := fmt.FormatString(s, v)
	h := r.hide()

	if _, ok := h.(fmt.Formatter); ok {
		panic("hide() must not return a value that implements fmt.Formatter")
	}

	fmt.Fprintf(s, f, h)
}
