package akin

import (
	"fmt"
	"reflect"
	"strings"
)

// A Value is a reflection-like representation of a Go value.
type Value struct {
	nat  any
	ref  reflect.Value
	expr ValueExpr
}

// valueOf returns the [Value] of v.
func valueOf(v any, e ValueExpr) Value {
	rvalue := reflect.ValueOf(v)

	if !rvalue.IsValid() {
		// If rvalue is invalid, that means v itself is the nil any interface.
		// So we represent that explicitly, rather than via an zero-valued
		// [reflect.Value].
		rvalue = reflect.ValueOf(&v).Elem()
	}

	return Value{v, rvalue, e}
}

// Expr returns the [ValueExpr] that describes how v was obtained.
func (v Value) Expr() ValueExpr {
	return v.expr
}

func valueFrom(v reflect.Value) Value {
	return Value{v.Interface(), v}
}

// Type returns the [Type] of v.
func (v Value) Type() Type {
	return Type{v.ref.Type()}
}

// isNil returns true if the value represented by v is nil.
func (v Value) isNil() bool {
	return v.Type().isNilable() && v.ref.IsNil()
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

	if v.ref.CanComplex() {
		s := fmt.Sprint(v.ref.Interface())
		return s[1 : len(s)-1]
	}

	if v.ref.CanFloat() {
		s := fmt.Sprint(v.ref.Interface())
		if strings.ContainsAny(s, ".eE") {
			return s
		}
		return s + ".0"
	}

	if v.ref.Kind() == reflect.String {
		return fmt.Sprintf("%q", v.ref.Interface())
	}

	return fmt.Sprintf("%v", v.ref.Interface())
}
