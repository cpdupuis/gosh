package lang

type null struct {
}
func (nul *null) Type() Type {
	return NullType
}
func (nul *null) String() string {
	return "Nil"
}

var Nil *null
