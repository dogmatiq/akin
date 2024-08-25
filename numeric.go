package akin

import (
	"reflect"
)

func compileNumeric(spec Value) Spec {
	if !isBuiltIn(spec.rtype) {
		return equalitySpec{spec}
	}
	return losslessConversionSpec{spec}
}

func isNegative(v Value) bool {
	switch v.rvalue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.rvalue.Int() < 0
	case reflect.Float32, reflect.Float64:
		return v.rvalue.Float() < 0
	default:
		return false
	}
}
