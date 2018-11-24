package lang_test

import (
	"github.com/cpdupuis/gosh/lang"
	"testing"
)

func TestStringType(t *testing.T) {
	s := &lang.String{}
	if s.Type() != lang.StringType {
		t.Fail()
	}
}

func TestStringCompare(t *testing.T) {
	s1 := lang.String{Str: "hello"}
	s2 := lang.String{Str: "hello"}
	if s1 != s2 {
		t.Fail()
	}
}

func TestStringString(t *testing.T) {
	s := lang.String{Str: "hi there"}
	if s.String() != "\"hi there\"" {
		t.Errorf("Expected quotes, got %v", s.String())
	}
}