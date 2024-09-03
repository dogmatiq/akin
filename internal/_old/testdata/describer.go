package akin

type describer struct {
	desc string
}

func (v *describer) Set(format string, args ...any) {
	v.desc = sprintf(format, args...)
}

func (v *describer) Const(p Const) {

}
