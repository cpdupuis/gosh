package lang

import (
	"regexp"
	"fmt"
	"strconv"
	"strings"
)

type ParseStatus int
const (
	OK = iota
	CloseSExp
)

func ParseSExp(inCh chan string) (Value, ParseStatus) {
	next := <-inCh
	if next == "null" {
		return Nil,OK
	}
	if next == ")" {
		return Nil,CloseSExp
	}
	if (next == "(") {
		// We're going to build a list until the closing )
		var res *Cons
		var curr *Cons
		for {
			item,status := ParseSExp(inCh)
			fmt.Printf("Here is item: %v\n", item)
			if status == CloseSExp {
				curr.Cdr = Nil
				return res, OK
			} else {
				newcons := &Cons{Car:item}
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
		return &Int{Number: intVal}, OK
	}
	match, err := regexp.MatchString("^\".*\"$", next)
	if match {
		return &String{Str: strings.Trim(next, "\"")}, OK
	}

	// Oh hey, maybe this is a symbol. Yeah, let's call it a symbol.
	return &Symbol{Sym: next}, OK
}
