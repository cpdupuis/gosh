package lang_test

import (
	"testing"
	"github.com/cpdupuis/gosh/lang"
)


func TestBuiltinPlus(t *testing.T) {
	scope := lang.NewScope(nil)
	lang.DefineIntBinOps(scope)
	cons := &lang.Cons{Car:&lang.Int{Number: 3}, Cdr: lang.Nil}
	cons = &lang.Cons{Car:&lang.Int{Number: 5}, Cdr: cons}
	cons = &lang.Cons{Car:&lang.Symbol{Sym:"+"}, Cdr: cons}
	ec := &lang.EvalContext{}
	ec.Push(lang.StandardForm)
	res,err := cons.Eval(scope, ec)
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
