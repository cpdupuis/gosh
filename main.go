package main

import (
	"fmt"
	"os"
	"github.com/cpdupuis/gosh/lang"
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


func treeize(inCh chan string, outCh chan lang.Value) lang.Value {
	for {
		sexp,_ := lang.ParseSExp(inCh)
		outCh <- sexp
	}
}

func main() {
	tokenCh := make(chan string, 256)
	treeCh := make(chan lang.Value, 256)
	go lang.Tokenize(os.Stdin, tokenCh)
	go treeize(tokenCh, treeCh)
	repl(treeCh)
}
