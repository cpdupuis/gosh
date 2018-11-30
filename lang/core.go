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

type Value interface {
	Type() Type
	String() string
	Eval(sc *Scope, ec *EvalContext) (Value,error)
}

type EvalContextFrame struct {
	CurrentLocation int
	ShouldEvalLocation func(int)
}

type EvalContext struct {
	Frames []*EvalContextFrame
}
