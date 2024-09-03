package akin

import "github.com/dogmatiq/akin/internal/fmtx"

// pick returns the string representation of t or f based on a truth value.
// Deprecated: dont
func pick[B ~bool, T any](truth B, t, f T) string {
	if truth {
		return fmtx.P(t)
	}
	return fmtx.P(f)
}
