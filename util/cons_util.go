package util

import (
	"github.com/cpdupuis/gosh/lang"
)

func listToSliceHelper(list lang.List, res []lang.Value) []lang.Value {
	res = append(res, list.First())
	// This only works if lists are only cons or nil
	if c,ok := list.Rest().(*lang.Cons); ok {
		return listToSliceHelper(c, res)
	} else {
		return res
	}
}

func ListToSlice(list lang.List) []lang.Value {
	var res []lang.Value
	return listToSliceHelper(list, res)
}

