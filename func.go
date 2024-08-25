package akin

import "reflect"

func compileFunc(spec reflect.Value) Spec {
	return specFunc(func(arg reflect.Value) error {
		if err := requireKind(reflect.Func, arg); err != nil {
			return err
		}
		// TODO
		return nil
	})
}
