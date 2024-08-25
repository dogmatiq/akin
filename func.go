package akin

import "reflect"

func compileFunc(Value) Spec {
	return specFunc(func(arg Value) error {
		if err := requireKind(reflect.Func, arg); err != nil {
			return err
		}
		// TODO
		return nil
	})
}
