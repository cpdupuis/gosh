package lang_test

import (
	"testing"
	"github.com/cpdupuis/gosh/lang"
)


func TestBuiltinPlus(t *testing.T) {
	scope := lang.NewScope(nil)
	lang.DefineIntBinOps(scope)
	cons := &lang.Cons{First:&lang.Int{Number: 3}, Rest: lang.Nil}
	cons = &lang.Cons{First:&lang.Int{Number: 5}, Rest: cons}
	cons = &lang.Cons{First:&lang.Symbol{Sym:"+"}, Rest: cons}
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
