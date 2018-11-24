package lang

type Null struct {

}
func (nul *Null) Type() Type {
	return NullType
}
func (nul *Null) String() string {
	return "Nil"
}

var Nil *Null
