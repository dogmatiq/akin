package akin

import "reflect"

// import (
// 	"reflect"
// )

// func compileNumeric(spec value) Set {
// 	if !isBuiltIn(spec.dynamic) {
// 		return equalitySpec{spec}
// 	}
// 	return Convertible{spec}
// }

func isNegative(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() < 0
	case reflect.Float32, reflect.Float64:
		return v.Float() < 0
	default:
		return false
	}
}
