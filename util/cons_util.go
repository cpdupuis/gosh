package util

import (
	"github.com/cpdupuis/gosh/lang"
)

func consToSliceHelper(cons *lang.Cons, res []lang.Value) []lang.Value {
	res = append(res, cons.First)
	// This only works if lists are only cons or nil
	if c,ok := cons.Rest.(*lang.Cons); ok {
		return consToSliceHelper(c, res)
	} else {
		return res
	}
}

func ConsToSlice(cons *lang.Cons) []lang.Value {
	var res []lang.Value
	return consToSliceHelper(cons, res)
}