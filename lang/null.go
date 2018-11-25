package lang

type null struct {
}
func (nul *null) Type() Type {
	return NullType
}
func (nul *null) String() string {
	return "Nil"
}
func (nul *null) Eval(sc *Scope) Value {
	// evaluating null returns itself
	return Nil
}
var Nil *null
