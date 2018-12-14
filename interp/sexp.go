package interp

import (
	"regexp"
	"fmt"
	"strconv"
	"strings"
	"github.com/cpdupuis/gosh/lang"
)

type ParseStatus int
const (
	OK = iota
	CloseSExp
)

func ParseSExp(inCh chan string) (lang.Value, ParseStatus) {
	next := <-inCh
	if next == "null" {
		return lang.Nil,OK
	}
	if next == ")" {
		return lang.Nil,CloseSExp
	}
	if (next == "(") {
		// We're going to build a list until the closing )
		var res *lang.Cons
		var curr *lang.Cons
		for {
			item,status := ParseSExp(inCh)
			fmt.Printf("Here is item: %v\n", item)
			if status == CloseSExp {
				curr.Cdr = lang.Nil
				return res, OK
			} else {
				newcons := &lang.Cons{Car:item}
				if res == nil {
					res = newcons
					curr = newcons
				} else {
					curr.Cdr = newcons
					curr = newcons
				}
			}
		}
	}
	intVal, err := strconv.ParseInt(next, 10, 64)
	if err == nil {
		// It's an int!
		return &lang.Int{Number: intVal}, OK
	}
	match, err := regexp.MatchString("^\".*\"$", next)
	if match {
		return &lang.String{Str: strings.Trim(next, "\"")}, OK
	}

	// Oh hey, maybe this is a symbol. Yeah, let's call it a symbol.
	return &lang.Symbol{Sym: next}, OK
}
