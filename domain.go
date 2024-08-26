package akin

import (
	"reflect"
)

// DomainFor returns the [Set] of all values of type T.
func DomainFor[T any]() Set {
	return domain{reflect.TypeFor[T]()}
}

// domainOf is the [Set] of all values of type T.
type domain struct {
	t reflect.Type
}

func (s domain) Contains(v any) bool {
	return reflect.TypeOf(v) == s.t
}

func (s domain) eval(v any) membership {
	if s.Contains(v) {
		return membership{
			IsMember: true,
			For:      []string{"it has type " + renderT(s.t)},
		}
	}

	return membership{
		IsMember: false,
		Against:  []string{"it does not have type " + renderT(s.t)},
	}
}

func (s domain) String() string {
	return "{ x | x inhabits " + renderT(s.t) + " }"
}
