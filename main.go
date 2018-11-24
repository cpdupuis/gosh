package main

import (
	"fmt"
	"github.com/cpdupuis/gosh/lang"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type handlerFunc func([]string) (string, error)

var handlerTable = map[string]handlerFunc{}

func def(words []string) (string, error) {
	return "", nil
}

func init() {
	handlerTable["def"] = def
}

func repl(treeCh chan lang.Value) {
	for {
		tree := <-treeCh
		fmt.Printf("Tree: <%+v>\n", tree)
	}
}

type ParseStatus int
const (
	OK = iota
	CloseSExp
	Dot
)

func parseSExp(inCh chan string) (lang.Value, ParseStatus) {
	next := <-inCh
	if next == "null" {
		return lang.Nil,OK
	}
	if next == ")" {
		return lang.Nil,CloseSExp
	}
	if next == "." {
		return lang.Nil, Dot
	}
	if (next == "(") {
		// We're going to build a list until the closing )
		var res *lang.Cons
		var curr *lang.Cons
		for {
			item,status := parseSExp(inCh)
			fmt.Printf("Here is item: %v\n", item)
			if status == CloseSExp {
				curr.Rest = lang.Nil
				return res, OK
			} else if status == Dot {
				curr.Rest, status = parseSExp(inCh)
				return res,OK
			} else {
				newcons := &lang.Cons{First:item}
				if res == nil {
					res = newcons
					curr = newcons
				} else {
					curr.Rest = newcons
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
	return lang.Nil, OK
}

func treeize(inCh chan string, outCh chan lang.Value) lang.Value {
	for {
		sexp,_ := parseSExp(inCh)
		outCh <- sexp
	}
}

func main() {
	hello := &lang.String{Str: "Hello"}
	world := &lang.String{Str: "world"}
	num := &lang.Int{Number: 123}
	a := &lang.Cons{First: world, Rest: lang.Nil}
	b := &lang.Cons{First: hello, Rest: a}
	fmt.Printf("My first cons: %v\n", b)
	d := &lang.Cons{First: hello, Rest: num}
	c := &lang.Cons{First: b, Rest: d}
	fmt.Printf("My second cons: %v\n", c)
	tokenCh := make(chan string, 256)
	treeCh := make(chan lang.Value, 256)
	go lang.Tokenize(os.Stdin, tokenCh)
	go treeize(tokenCh, treeCh)
	repl(treeCh)
}
