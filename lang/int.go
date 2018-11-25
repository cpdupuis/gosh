package lang

import (
	"strconv"
)

type Int struct {
	Number int64
}
func (i *Int) Type() Type {
	return IntType
}
func (i *Int) String() string {
	return strconv.FormatInt(i.Number, 10)
}

func (i *Int) Eval(sc *Scope) Value {
	// evaluating an int returns itself
	return i
}