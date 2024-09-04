package akin

import (
	"fmt"
	"strings"
)

func stringP(p Predicate, f form) string {
	var w strings.Builder
	p.visitP(&stringer{f, &w})
	return w.String()
}

func stringA(a Attribute, f form) string {
	var w strings.Builder
	a.visitA(&stringer{f, &w})
	return w.String()
}

func stringR(r Rationale) string {
	var w strings.Builder
	r.visitR(&stringer{affirmative, &w})
	return w.String()
}

type (
	// form represents the way avalue that is rendered. It is rendered either in
	// the [affirmative] or [negative] form.
	form bool

	renderer interface {
		form() form
		buf() *strings.Builder
	}

	stringer struct {
		f form
		b *strings.Builder
	}
)

const (
	// affirmative form is the normal form of a [Predicate] or [Attribute].
	// For example, 𝒙 ≡ 𝒆 is the affirmative form of the [Identity] predicate.
	affirmative form = true

	// negative form is the negative form of a [Predicate] or [Attribute]. For
	// example, 𝒙 ≢ 𝒆 is the negative form of the [Identity] predicate.
	negative form = false
)

func (s *stringer) form() form {
	return s.f
}

func (s *stringer) buf() *strings.Builder {
	return s.b
}

// render appends a formatted string to the stringer's output.
//
// Portions of the format string contained in braces are replaced according
// [stringer.Form]. For example, if the format string is "𝒙 {≡|≢} 𝒆", the
// output will be "𝒙 ≡ 𝒆" when the form is [canonical], or "𝒙 ≢ 𝒆" when the
// form is [negated].
func render(s renderer, format string, args ...any) {
	renderNegatable(s, affirmative, format, args...)
}

// renderNegatable is a specialization of [render] that inverts the rendering
// behavior if isCanonical is false.
func renderNegatable[B ~bool](
	r renderer,
	isCanonical B,
	format string,
	args ...any,
) {
	var w strings.Builder

	f := r.form()
	if !isCanonical {
		f = !f
	}

	for {
		begin := strings.IndexRune(format, '{')
		if begin == -1 {
			break
		}
		segment := format[begin+1:]

		end := strings.IndexRune(segment, '}')
		if end == -1 {
			break
		}
		segment = segment[:end]

		w.WriteString(format[:begin])
		format = format[begin+end+2:]

		pipe := strings.IndexRune(segment, '|')

		if f == negative {
			w.WriteString(segment[pipe+1:])
		} else if pipe != -1 {
			w.WriteString(segment[:pipe])
		}
	}

	w.WriteString(format)
	fmt.Fprintf(r.buf(), w.String(), args...)
}

// subscript renders n as a string using unicode subscript characters.
func subscript[
	T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64,
](n T) string {
	digits := []rune(fmt.Sprint(n))

	for i, r := range digits {
		switch r {
		case '-':
			digits[i] = '₋'
		default:
			digits[i] = r - '0' + '₀'
		}
	}

	return string(digits)
}

// parens adds mathematical parentheses to an expression if it contains spaces
// (and does not already have them).
func parens(s string) string {
	if strings.HasPrefix(s, "❨") {
		return s
	}
	if !strings.ContainsAny(s, " ") {
		return s
	}
	return "❨" + s + "❩"
}
