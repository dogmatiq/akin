// fn is a [Set] implemented as an eval function.
type fn func(any) error

func (s fn) Contains(v any) bool {
	return s.Eval(v).IsMember
}

func (s fn) Eval(v any) Membership {
	if err := s(v); err != nil {
		return Membership{IsMember: false, Reason: err.Error()}

	}
	return Membership{
		IsMember: true,
	}
}
