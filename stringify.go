package akin

import (
	"fmt"
	"strings"
)

// form represents the "form" of some value that is rendered. That is, whether
// or not it is inverted.
type form bool

const (
	canonical form = true
	negated   form = false
)

func stringP(p Predicate, f form) string {
	s := stringer{Form: f}
	p.VisitP(&s)
	return s.Output
}

func stringA(a Attribute, f form) string {
	s := stringer{Form: f}
	a.VisitA(&s)
	return s.Output
}

func stringR(r Rationale) string {
	s := stringer{Form: canonical}
	r.VisitR(&s)
	return s.Output
}

type stringer struct {
	Form   form
	Output string
}

// write appends a formatted string to the stringer's output.
//
// Portions of the format string contained in braces are replaced according
// [stringer.Form]. For example, if the format string is "𝒙 {≡|≢} 𝒆", the
// output will be "𝒙 ≡ 𝒆" when the form is [canonical], or "𝒙 ≢ 𝒆" when the
// form is [negated].
func write(s *stringer, format string, args ...any) {
	writeNegatable(s, canonical, format, args...)
}

// writeNegatable is a specialization of [write] that inverts the rendering
// behavior if isCanonical is false.
func writeNegatable[T ~bool](s *stringer, isCanonical T, format string, args ...any) {
	var w strings.Builder

	f := s.Form
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

		if f == negated {
			w.WriteString(segment[pipe+1:])
		} else if pipe != -1 {
			w.WriteString(segment[:pipe])
		}
	}

	w.WriteString(format)
	s.Output += fmt.Sprintf(w.String(), args...)
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
