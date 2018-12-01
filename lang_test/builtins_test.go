package lang_test

import (
	"testing"
	"github.com/cpdupuis/gosh/lang"
)

func TestBuiltinPlus(t *testing.T) {
	paramNames := []string{"foo", "bar"}
	lambda := lang.CreateBuiltin(paramNames, lang.BuiltinPlus, lang.StandardForm)
	scope := lang.NewScope(nil)
	cons := &lang.Cons{First:&lang.Int{Number: 3}, Rest: lang.Nil}
	cons = &lang.Cons{First:&lang.Int{Number: 5}, Rest: cons}
	ec := &lang.EvalContext{}
	ec.Push(lang.StandardForm)
	res,err := lambda.Call(scope, ec, cons)
	if err != nil {
		t.Fail()
	}
	if num,ok := res.(*lang.Int); ok {
		if num.Number != 8 {
			t.Errorf("Wrong number: %+v", num)
		}
	} else {
		t.Fail()
	}
}

func TestBuiltinCons(t *testing.T) {
	paramNames := []string{"foo", "bar"}
	lambda := lang.CreateBuiltin(paramNames, lang.BuiltinCons, lang.StandardForm)
	scope := lang.NewScope(nil)
	ec := &lang.EvalContext{}
	ec.Push(lang.StandardForm)
	cons := &lang.Cons{First:&lang.Int{Number: 3}, Rest: lang.Nil}
	cons = &lang.Cons{First:&lang.Int{Number: 5}, Rest: cons}

	seven := &lang.Int{Number: 7}

	consArgs := &lang.Cons{First:cons, Rest: lang.Nil}
	consArgs = &lang.Cons{First:seven, Rest: consArgs}

	res,err := lambda.Call(scope, ec, consArgs)
	if err != nil {
		t.Errorf("Failed: %s", err.Error())
	}
	if c,ok := res.(*lang.Cons); ok {
		str := c.String()
		if str != "( 7 5 3 )" {
			t.Errorf("Wrong list: %+v", str)
		}
	} else {
		t.Errorf("Res is not cons: %+v", res)
	}
}
