package lang_test

import (
	"testing"
	"github.com/cpdupuis/gosh/lang"
)

func TestPush(t *testing.T) {
	ec := &lang.EvalContext{}
	ec.Push(lang.StandardForm)
	if len(ec.Frames) != 1 {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	ec := &lang.EvalContext{}
	ec.Push(lang.StandardForm)
	ec.Pop()
	if len(ec.Frames) != 0 {
		t.Fail()
	}
}


func TestPopEmpty(t *testing.T) {
	ok := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				// recovered from panic
				ok = true
			}
		}()
		ec := &lang.EvalContext{}
		ec.Pop()	
	}()
	if !ok {
		t.Fail()
	}		
}

func TestPushityPoppityPop(t *testing.T) {
	ec := &lang.EvalContext{}
	for i:=0; i<1024; i++ {
		ec.Push(lang.StandardForm)
	}
	for i:=0; i<1024; i++ {
		ec.Pop()
	}
	ok := false
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				// recovered from panic
				ok = true
			}
		}()
		ec.Pop()	
	}()
	if !ok {
		t.Fail()
	}		
}
