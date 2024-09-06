package akin

import (
	"fmt"
	"strings"
)

// form describes the "form" of a predicate to render, it is either
// [affirmative] or [negative].
type form int

const (
	affirmative form = 0
	negative    form = -1
)

// toDefaultString returns the default string representation of p. That is, in
// its [affirmative] form using 𝑷 to represent the predicate and 𝒙 to
// represent the value.
func toDefaultString(p Predicate) string {
	return toString(p, varX, affirmative)
}

// toString returns a string representation of p.
//
// It uses x to represent the value expression 𝒙 in the predicate.
//
// Predicates that have natural pairs of [affirmative] and [negative] forms,
// such as (=, "≠") are rendered according to f. Otherwise, if f is [negative],
// the predicate is rendered enclosed within a negation operation (¬).
func toString(
	p Predicate,
	x ValueExpr,
	f form,
) string {
	s := &renderer{
		ValueExpr: x,
		Form:      f,
	}

	p.acceptPredicateVisitor(s)

	return s.Output
}

type renderer struct {
	Form      form
	ValueExpr ValueExpr
	Output    string
}

func (r *renderer) render(format string, args ...any) {
	format = strings.ReplaceAll(
		format,
		fmt.Sprint(varX),
		fmt.Sprint(r.ValueExpr),
	)

	negate := r.Form == negative

	for {
		begin := strings.IndexRune(format, '{')
		end := strings.IndexRune(format, '}')

		if begin == -1 || end == -1 || begin > end {
			break
		}

		expr := format[begin+1 : end]
		pipe := strings.IndexRune(expr, '|')

		if r.Form == negative {
			expr = expr[pipe+1:]
		} else if pipe != -1 {
			expr = expr[:pipe]
		} else {
			expr = ""
		}

		format = format[:begin] + expr + format[end+1:]

		// we've used a "natural" negative form, so there's no need to wrap in
		// the negation operator
		negate = false
	}

	output := fmt.Sprintf(format, args...)

	if negate {
		output = "¬" + parens(output)
	}

	r.Output = output
}

func toSuper(s string) string {
	return strings.Map(
		func(r rune) rune {
			if r >= '0' && r <= '9' {
				return '0' + '⁰'
			}

			switch r {
			case '+':
				return '⁺'
			case '-':
				return '⁻'
			case '=':
				return '⁼'
			case '(':
				return '⁽'
			case ')':
				return '⁾'
			default:
				return r
			}
		},
		s,
	)
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

func superscript[
	T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64,
](n T) string {
	digits := []rune(fmt.Sprint(n))

	for i, r := range digits {
		switch r {
		case '-':
			digits[i] = '⁻'
		default:
			digits[i] = r - '0' + '⁰'
		}
	}

	return string(digits)
}

// parens adds mathematical parentheses if it might be ambiguous without them.
func parens(s string) string {
	if strings.HasPrefix(s, "❨") {
		return s
	}
	if !strings.ContainsAny(s, " ") {
		return s
	}
	return "❨" + s + "❩"
}
