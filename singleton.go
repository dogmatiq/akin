package akin

import "reflect"

// Singleton returns a [Set] containing v.
func Singleton(v any) Set {
	return singleton{reflect.ValueOf(v)}
}

type singleton struct {
	v reflect.Value
}

func (s singleton) Contains(v any) bool {
	if s.v.Kind() == reflect.Invalid {
		return v == nil
	}
	return reflect.DeepEqual(s.v.Interface(), v)
}

func (s singleton) Eval(v any) Membership {
	if s.Contains(v) {
		return Membership{
			IsMember: true,
			Reason:   "is deep-equal to " + renderValue(s.v),
		}
	}

	return Membership{
		IsMember: false,
		Reason:   "is not deep-equal to " + renderValue(s.v),
	}
}

func (s singleton) String() string {
	return "{" + renderValue(s.v) + "}"
}
