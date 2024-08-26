package akin

import "reflect"

// value is a potential member of a [Set].
type value struct {
	v       any
	r       reflect.Value
	dynamic reflect.Type
	static  reflect.Type
}

// valueOf returns a new [value] representing v.
func valueOf[T any](v T) value {
	s := reflect.TypeFor[T]()
	d := reflect.TypeOf(v)
	if d == nil {
		d = s
	}

	return value{
		v:       v,
		r:       reflect.ValueOf(v),
		dynamic: d,
		static:  s,
	}
}
