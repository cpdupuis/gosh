package main

import (
	"fmt"
	"os"
	"github.com/cpdupuis/gosh/lang"
)

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
