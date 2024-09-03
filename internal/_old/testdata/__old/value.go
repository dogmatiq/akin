package akin

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/dogmatiq/akin/internal/reflectx"
)

// ValueEquivalence is a [Property] that is true when the [Value] is equal to
// some specific, but abstract value.
//
// Abstract in this context means that the value is conceptual, such as the
// number one, versus a specific Go value, such as float64(1.0).
type ValueEquivalence struct {
	V string
}

func (p ValueEquivalence) visitA(v PropertyVisitor) { v.ValueEquivalence(p) }
func (p ValueEquivalence) String() string           { return "𝒙 ≍ " + p.V }
func (p ValueEquivalence) NString() string          { return "𝒙 ≭ " + p.V } //revive:disable-line:exported

// A Value is a Go value against which a [Predicate] may be evaluated.
type Value struct {
	rvalue reflect.Value
}

// Type returns the Go type of the value.
func (x Value) Type() Type {
	return Type{x.rvalue.Type()}
}

func (x Value) String() string {
	switch x.rvalue.Type() {
	case reflect.TypeOf(""),
		reflect.TypeOf(true),
		reflect.TypeOf(0),
		reflect.TypeOf(0.0),
		reflect.TypeOf(0 + 0i):
		return x.valuePart()
	}

	return x.Type().String() + "(" + x.valuePart() + ")"
}

func (x Value) valuePart() string {
	if reflectx.IsNil(x.rvalue) {
		return "nil"
	}

	if x.rvalue.CanComplex() {
		s := fmt.Sprint(x.rvalue.Interface())
		return s[1 : len(s)-1]
	}

	if x.rvalue.CanFloat() {
		s := fmt.Sprint(x.rvalue.Interface())
		if strings.ContainsAny(s, ".eE") {
			return s
		}
		return s + ".0"
	}

	if x.rvalue.Kind() == reflect.String {
		return fmt.Sprintf("%q", x.rvalue.Interface())
	}

	return fmt.Sprintf("%v", x.rvalue.Interface())
}

func (x Value) isNil() bool {
	return x.Type().isNilable() && x.rvalue.IsNil()
}
