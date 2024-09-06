package akin

import (
	"fmt"
	"strings"
)

// logicalForm describes the "logicalForm" of a predicate to render, it is
// either [affirmativeForm] or [negativeForm].
type logicalForm int

const (
	affirmativeForm logicalForm = iota
	negativeForm
)

func predicateToString(p Predicate) string {
	var w strings.Builder

	p.acceptPredicateVisitor(&predicateRenderer{
		PredicateExpr: defaultPredicateExpr,
		ValueExpr:     defaultValueExpr,
		Form:          affirmativeForm,
		Output:        &w,
	})

	return w.String()
}

func rationaleToString(r Rationale) string {
	var w strings.Builder

	r.acceptRationaleVisitor(&rationaleRenderer{
		Output: &w,
	})

	return w.String()
}

type (
	predicateRenderer struct {
		PredicateExpr Expr
		ValueExpr     Expr
		Form          logicalForm
		Parenthesize  bool
		Output        *strings.Builder
	}

	rationaleRenderer struct {
		Output *strings.Builder
	}
)

func (r *predicateRenderer) Render(format string, args ...any) {
	fmt.Fprintf(r.Output, r.subst(format), args...)
}

func (r *predicateRenderer) subst(s string) string {
	negate := r.Form == negativeForm

	for {
		begin := strings.IndexRune(s, '{')
		end := strings.IndexRune(s, '}')

		if begin == -1 || end == -1 || begin > end {
			break
		}

		expr := s[begin+1 : end]
		pipe := strings.IndexRune(expr, '|')

		if r.Form == negativeForm {
			expr = expr[pipe+1:]
		} else if pipe != -1 {
			expr = expr[:pipe]
		} else {
			expr = ""
		}

		s = s[:begin] + expr + s[end+1:]

		// we've used a "natural" negative form, so there's no need to wrap in
		// the negation operator
		negate = false
	}

	s = strings.ReplaceAll(
		s,
		defaultPredicateExpr.String(),
		r.PredicateExpr.String(),
	)

	s = strings.ReplaceAll(
		s,
		defaultValueExpr.String(),
		r.ValueExpr.String(),
	)

	if r.Parenthesize {
		s = parens(s)
	}

	if negate {
		s = "¬" + s
	}

	return s
}

func (r *rationaleRenderer) Render(format string, args ...any) {
	fmt.Fprintf(r.Output, format, args...)
}

func super(format string, args ...any) string {
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
		fmt.Sprintf(format, args...),
	)
}

func sub(format string, args ...any) string {
	return strings.Map(
		func(r rune) rune {
			if r >= '0' && r <= '9' {
				return '0' + '₀'
			}

			switch r {
			case '+':
				return '₊'
			case '-':
				return '₋'
			case '=':
				return '₌'
			case '(':
				return '₍'
			case ')':
				return '₎'
			default:
				return r
			}
		},
		fmt.Sprintf(format, args...),
	)
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
