package main

import (
	"regexp"
	"strconv"
	"os"
	"fmt"
	"strings"
	"github.com/cpdupuis/gosh/lang"
)

type handlerFunc func([]string) (string,error)

var handlerTable = map[string]handlerFunc{}

func def(words []string) (string,error) {
	return "",nil
}

func init() {
	handlerTable["def"] = def
}

func repl(treeCh chan lang.Value) {
	for {
		tree := <- treeCh
		fmt.Printf("Tree: <%+v>\n", tree)
	}
}



func treeizeHelper(inCh chan string, curr lang.Value) lang.Value {
	next := <- inCh
	if next == "null" {
		return lang.Nil
	}
	if next == "(" {
		// ?
	}
	intVal,err := strconv.ParseInt(next, 10, 64)
	if err == nil {
		// It's an int!
		return &lang.Int{Number: intVal}
	}
	match, err := regexp.MatchString("^\".*\"$", next)
	if match {
		return &lang.String{Str: strings.Trim(next, "\"")}
	}
	return lang.Nil
}

func treeize(inCh chan string, outCh chan lang.Value) lang.Value {
	for {
		outCh <- treeizeHelper(inCh, lang.Nil)
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
