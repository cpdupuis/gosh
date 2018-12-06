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
	a := &lang.Cons{First: world, Rest: lang.Nil}
	b := &lang.Cons{First: hello, Rest: a}
	c := &lang.Cons{First: num, Rest: b}

	res := util.ConsToSlice(c)
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
