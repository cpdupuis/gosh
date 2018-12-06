package lang_test

import (
	"github.com/cpdupuis/gosh/lang"
	"testing"
)

func TestConsString(t *testing.T) {
	hello := &lang.String{Str: "Hello"}
	world := &lang.String{Str: "world"}
	num := &lang.Int{Number: 123}
	a := &lang.Cons{Car: world, Cdr: lang.Nil}
	b := &lang.Cons{Car: hello, Cdr: a}
	if b.String() != "( \"Hello\" \"world\" )" {
		t.Errorf("Wrong first string: %v\n", b.String())
	}
	e := &lang.Cons{Car:num, Cdr: lang.Nil}
	d := &lang.Cons{Car: hello, Cdr: e}
	c := &lang.Cons{Car: b, Cdr: d}
	if c.String() != "( ( \"Hello\" \"world\" ) \"Hello\" 123 )" {
		t.Errorf("Wrong second string: %v\n", c.String())
	}
}

