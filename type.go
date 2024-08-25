package akin

import "reflect"

func is[T any](t reflect.Type) bool {
	return t == reflect.TypeFor[T]()
}

func isBuiltIn(t reflect.Type) bool {
	return t.PkgPath() == "" && t.Name() != ""
}
