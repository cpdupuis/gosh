package lang_test

import (
	"testing"
	"github.com/cpdupuis/gosh/lang"
)

func TestBinaryPlus(t *testing.T) {
	paramNames := []string{"foo", "bar"}
	lambda := lang.CreateBuiltin(paramNames, lang.BuiltinPlus)
	scope := lang.NewScope(nil)
	cons := &lang.Cons{First:&lang.Int{Number: 3}, Rest: lang.Nil}
	cons = &lang.Cons{First:&lang.Int{Number: 5}, Rest: cons}
	res,err := lambda.Call(scope, cons)
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
