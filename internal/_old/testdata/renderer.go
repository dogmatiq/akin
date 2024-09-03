package akin

// import (
// 	"strings"
// )

// type (
// 	renderer struct {
// 		result  string
// 		inverse bool
// 		verbose bool
// 	}
// )

// func verboseA(a Property, pos bool) string {
// }

// // predicates ...
// func (v *renderer) Const(p Const) {
// 	if v.pos {
// 		v.render("{⊤|⊥}")
// 	} else {
// 		v.render("{⊥|⊤}")
// 	}
// }
// func (v *vrenderer) Const(p Const) { v.renderx(p == Top, "𝑷 is {satisfied|violated} by any 𝒙") }

// // properties ...
// func (v *renderer) HasValue(a HasValue)  { v.render("𝒙 {≍|≭} %s", a.Desc) }
// func (v *vrenderer) HasValue(a HasValue) { v.render("𝒙 is {not }equivalent to %s", a.Desc) }

// // reasons ...
// func (v *renderer) PConst(PConst)     { v.render("𝑷 is satisfied for any 𝒙") }
// func (v *renderer) PVacuous(PVacuous) { v.render("𝑷 has no constituents") }
// func (v *renderer) QTrue(r QTrue)     { v.render("%s holds, because %s", r.EQ.P, r.EQ.Reason) }
// func (v *renderer) QFalse(r QFalse)   { v.render("%s does not hold, because %s", r.EQ.P, r.EQ.Reason) }
// func (v *renderer) ATrue(r ATrue)     { v.result = verbose(r.A, true) }
// func (v *renderer) AFalse(r AFalse)   { v.result = verbose(r.A, false) }

// type (
// 	renderer struct {
// 		result string
// 		pos    bool
// 	}
// 	vrenderer struct{ renderer }
// )

// func (v *renderer) render(format string, args ...any) {
// 	var f strings.Builder

// 	for {
// 		begin := strings.IndexRune(format, '{')
// 		if begin == -1 {
// 			f.WriteString(format)
// 			break
// 		}
// 		segment := format[begin+1:]

// 		end := strings.IndexRune(segment, '}')
// 		if end == -1 {
// 			f.WriteString(format)
// 			break
// 		}
// 		segment = segment[:end]

// 		f.WriteString(format[:begin])
// 		format = format[end+1:]

// 		pipe := strings.IndexRune(segment, '|')

// 		if !v.pos {
// 			f.WriteString(segment[pipe+1:])
// 		} else if pipe != -1 {
// 			f.WriteString(segment[:pipe])
// 		}
// 	}

// 	v.result = sprintf(f.String(), args...)
// }

// func (v *renderer) renderx(pos bool, format string, args ...any) {
// 	if !pos {
// 		v.pos = !v.pos
// 	}

// 	v.render(format, args...)

// 	if !pos {
// 		v.pos = !v.pos
// 	}
// }
