package lang

type List interface {
	Type() Type
	String() string
	Eval(sc *Scope) Value
	Length() int // May need to be ListLength if we don't want Strings to be Lists...
}
