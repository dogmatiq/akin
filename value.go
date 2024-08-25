package akin

import "reflect"

// Value is a value that can be compared against a [Spec].
type Value struct {
	value  any
	rvalue reflect.Value
	rtype  reflect.Type
}

// ValueOf returns a new [Value] representing v.
func ValueOf[T any](v T) Value {
	return Value{
		value:  v,
		rvalue: reflect.ValueOf(v),
		rtype:  reflect.TypeOf(v),
	}
}
