package testx

// Err is a type that implements the [error] interface.
type Err struct{}

func (Err) Error() string {
	return "<error>"
}
