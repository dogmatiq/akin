package akin

// Optional represents an optional value of type T.
type Optional[T any] struct {
	v *T
}

// Some returns an Optional[T] containing v.
func Some[T any](v T) Optional[T] {
	return Optional[T]{&v}
}

// None returns an empty Optional[T].
func None[T any]() Optional[T] {
	return Optional[T]{}
}

// Get returns the value of o, if it has one; otherwise it panics.
func (o Optional[T]) Get() T {
	if o.v == nil {
		panic("value is not present")
	}
	return *o.v
}

// TryGet returns the value of o, if it has one; otherwise ok is false and v is
// the zero value of T.
func (o Optional[T]) TryGet() (v T, ok bool) {
	if o.v != nil {
		v = *o.v
		ok = true
	}
	return
}

// HasValue returns true if o has a value.
func (o Optional[T]) HasValue() bool {
	return o.v != nil
}
