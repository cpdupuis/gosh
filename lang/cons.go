package lang

import (
	"errors"
	"fmt"
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
	rest := cons.Rest
	firstVal,err := first.Eval(sc)
	if err != nil {
		return Nil,err
	}
	if lambda,ok := firstVal.(*Lambda); ok {
		return lambda.Call(sc, rest)
	} else {
		return Nil,errors.New(fmt.Sprintf("first is not a lambda: %+v", first))
	}
}

func (cons *Cons) Length() int {
	list,ok := cons.Rest.(List)
	if ok {
		return 1 + list.Length()
	} else {
		return 2 // Why 2? Because rest is not a Cons and not Nil. So I guess my length is 2.
	}
}
