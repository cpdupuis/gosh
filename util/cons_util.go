package util

import (
	"github.com/cpdupuis/gosh/lang"
)

func listToSliceHelper(list lang.List, res []lang.Value) []lang.Value {
	res = append(res, list.First)
	// This only works if lists are only cons or nil
	if c,ok := cons.Rest.(*lang.Cons); ok {
		return consToSliceHelper(c, res)
	} else {
		return res
	}
}

func ListToSlice(list lang.List) []lang.Value {
	var res []lang.Value
	return listToSliceHelper(list, res)
}

