package lang_test

import (
	"testing"
	"github.com/cpdupuis/gosh/lang"
)

func TestPush(t *testing.T) {
	ec := &lang.EvalContext{}
	ec.Push()
	if len(ec.Frames) != 1 {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	ec := &lang.EvalContext{}
	ec.Push()
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

