package lang

type Type int
const (
	SymbolType = iota
	StringType
	ConsType
	IntType
	LambdaType
	NullType
	ScopeType
)

type Value interface{
	Type() Type
	String() string
	Eval(sc *Scope) Value
}
