package lang

type Type int
const (
	SymbolType = iota
	StringType
	ConsType
	IntType
	NullType
	ScopeType
)

type Value interface{
	Type() Type
	String() string
}
