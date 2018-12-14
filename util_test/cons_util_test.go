package util_test

import (
	"github.com/cpdupuis/gosh/lang"
	"github.com/cpdupuis/gosh/util"
	"testing"
)

func TestConsToSlice(t *testing.T) {
	hello := &lang.String{Str: "Hello"}
	world := &lang.String{Str: "world"}
	num := &lang.Int{Number: 123}
	a := &lang.Cons{Car: world, Cdr: lang.Nil}
	b := &lang.Cons{Car: hello, Cdr: a}
	c := &lang.Cons{Car: num, Cdr: b}

	res := util.ListToSlice(c)
	if res[0].String() != "123" {
		t.Fail()
	}
	if res[1].String() != "\"Hello\"" {
		t.Fail()
	}
	if res[2].String() != "\"world\"" {
		t.Fail()
	}
}
