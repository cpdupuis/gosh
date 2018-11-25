package lang

import (
	"strings"
)

type Cons struct {
	First Value
	Rest Value
}

func (cons *Cons) Type() Type {
	return ConsType
}

func (cons *Cons) String() string {
	strarray := make([]string, 0)
	strarray = append(strarray, "(")
	strarray = append(strarray, cons.First.String())
	curr := cons.Rest

Loop:	
	for {
		switch item := curr.(type) {
		case *null:
			break Loop
		case *Cons:
			strarray = append(strarray, item.First.String())
			curr = item.Rest
		default:
			strarray = append(strarray, ".")
			strarray = append(strarray, curr.String())
			break Loop
		}
	}
	strarray = append(strarray, ")")
	return strings.Join(strarray, " ")
}

func (cons *Cons) Eval(sc *Scope) Value {
	return Nil
}
