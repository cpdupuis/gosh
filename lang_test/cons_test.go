package lang_test

import (
	"github.com/cpdupuis/gosh/lang"
	"testing"
)

func TestConsString(t *testing.T) {
	hello := &lang.String{Str: "Hello"}
	world := &lang.String{Str: "world"}
	num := &lang.Int{Number: 123}
	a := &lang.Cons{First: world, Rest: lang.Nil}
	b := &lang.Cons{First: hello, Rest: a}
	if b.String() != "( \"Hello\" \"world\" )" {
		t.Errorf("Wrong first string: %v\n", b.String())
	}
	e := &lang.Cons{First:num, Rest: lang.Nil}
	d := &lang.Cons{First: hello, Rest: e}
	c := &lang.Cons{First: b, Rest: d}
	if c.String() != "( ( \"Hello\" \"world\" ) \"Hello\" 123 )" {
		t.Errorf("Wrong second string: %v\n", c.String())
	}
}

func TestConsToSlice(t *testing.T) {
	hello := &lang.String{Str: "Hello"}
	world := &lang.String{Str: "world"}
	num := &lang.Int{Number: 123}
	a := &lang.Cons{First: world, Rest: lang.Nil}
	b := &lang.Cons{First: hello, Rest: a}
	c := &lang.Cons{First: num, Rest: b}

	res := c.ToSlice()
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