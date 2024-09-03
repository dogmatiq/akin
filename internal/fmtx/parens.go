package fmtx

import "strings"

// Parens adds mathematical parentheses to an expression if it contains spaces
// (and does not already have them).
func Parens(v any) string {
	s := P(v)

	if strings.HasPrefix(s, "❨") {
		return s
	}

	if !strings.ContainsAny(s, " ") {
		return s
	}

	return "❨" + s + "❩"
}
