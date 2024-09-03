package akin

import (
	"strings"
)

// func render(v any) string {
// 	switch v := v.(type) {
// 	case reflect.Type:
// 		return formatT(v)
// 	case reflect.Value:
// 		return formatTV(v)
// 	}
// 	return formatTV(reflectx.ValueOf(v))
// }

func parens(s string) string {
	if strings.HasPrefix(s, "❨") {
		return s
	}

	if !strings.ContainsAny(s, " ") {
		return s
	}

	return "❨" + s + "❩"
}

func parensf(format string, args ...any) string {
	s := renderf(format, args...)
	return parens(s)
}

// join renders a list of elements as a string, with each element
// separated by a separator string.
func join[T any](sep string, elements ...T) string {
	return join2(sep, sep, elements...)
}

// join2 renders a list of elements as a string, with each element
// separated by a separator string, and a different separator string
// before the last element.
func join2[T any](sep, lastSep string, elements ...T) string {
	var w strings.Builder
	last := len(elements) - 1

	for pos, elem := range elements {
		switch pos {
		case 0:
			// never add a separator before the first element, even if it's also
			// the last element.
		case last:
			w.WriteString(lastSep)
		default:
			w.WriteString(sep)
		}

		w.WriteString(render(elem))
	}

	return w.String()
}

func choose[B ~bool, T any](cond B, t, f T) T {
	if cond {
		return t
	}
	return f
}

func insert[B ~bool](cond B, s any) string {
	if cond {
		return renderf(" %s", s)
	}
	return ""
}
