package lang

import (
	"strings"
)

type Cons struct {
	First Value
	Rest List
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
			panic("Unexpected!")
		}
	}
	strarray = append(strarray, ")")
	return strings.Join(strarray, " ")
}

func (cons *Cons) Eval(sc *Scope) (Value,error) {
	// Eval'ing a cons means calling the lambda in the car with the cons in the cdr.
	first := cons.First
	firstVal,err := first.Eval(sc)
	if err != nil {
		return Nil,err
	}
	if lambda,ok := firstVal.(*Lambda); ok {
		return lambda.Call(sc, cons.Rest)
	} else {
		return cons,nil
	}
}

func (cons *Cons) Length() int {
	return 1 + cons.Rest.Length()
}

