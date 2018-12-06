package lang

type List interface {
	First() Value
	Rest() List
	Type() Type
	String() string
	Eval(sc *Scope, ec *EvalContext) (Value,error)
	Length() int // May need to be ListLength if we don't want Strings to be Lists...
}
