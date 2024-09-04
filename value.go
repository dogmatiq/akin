package akin

import (
	"fmt"
	"reflect"
	"strings"
)

// A Value is a reflection-like representation of a Go value.
type Value struct {
	rvalue reflect.Value
}

// valueOf returns the [Value] of v.
func valueOf(v any) Value {
	rvalue := reflect.ValueOf(v)

	if !rvalue.IsValid() {
		// If rvalue is invalid, that means v itself is the nil any interface.
		// So we represent that explicitly, rather than via an zero-valued
		// [reflect.Value].
		rvalue = reflect.ValueOf(&v).Elem()
	}

	return Value{rvalue}
}

// Type returns the [Type] of v.
func (v Value) Type() Type {
	return Type{v.rvalue.Type()}
}

// isNil returns true if the value represented by v is nil.
func (v Value) isNil() bool {
	return v.Type().isNilable() && v.rvalue.IsNil()
}

func (v Value) String() string {
	if v.Type().isDefaultForConstant() {
		return v.valueString()
	}
	return v.Type().String() + "(" + v.valueString() + ")"
}

func (v Value) valueString() string {
	if v.isNil() {
		return "nil"
	}

	if v.rvalue.CanComplex() {
		s := fmt.Sprint(v.rvalue.Interface())
		return s[1 : len(s)-1]
	}

	if v.rvalue.CanFloat() {
		s := fmt.Sprint(v.rvalue.Interface())
		if strings.ContainsAny(s, ".eE") {
			return s
		}
		return s + ".0"
	}

	if v.rvalue.Kind() == reflect.String {
		return fmt.Sprintf("%q", v.rvalue.Interface())
	}

	return fmt.Sprintf("%v", v.rvalue.Interface())
}
