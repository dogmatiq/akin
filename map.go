package akin

import "reflect"

func compileMap(Value) Spec {
	return specFunc(func(arg Value) error {
		if err := requireKind(reflect.Map, arg); err != nil {
			return err
		}
		// TODO
		return nil
	})
}
