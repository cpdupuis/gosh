package main

import (
	"strconv"
	"io"
	"os"
	"fmt"
	"bufio"
	"strings"
)


type Type int
const (
	SymbolType = iota
	StringType
	ConsType
	IntType
	NullType
	ScopeType
)

type Value interface{
	Type() Type
	String() string
}

type Cons struct {
	First *Value
	Rest *Value
}

func (cons *Cons) Type() Type {
	return ConsType
}
func (cons *Cons) String() string {
	return "cons"
}

type Scope struct {
	ScopeTable map[string]Value
}

func (scope *Scope) Type() Type {
	return ScopeType
}

func (scope *Scope) String() string {
	return "scope"
}

type String struct {
	Str string
}
func (str *String) Type() Type {
	return StringType
}
func (str *String) String() string {
	return str.Str
}
type Int struct {
	Number int64
}
func (i *Int) Type() Type {
	return IntType
}
func (i *Int) String() string {
	return strconv.FormatInt(i.Number, 10)
}

type Null struct {

}
func (nul *Null) Type() Type {
	return NullType
}
func (nul *Null) String() string {
	return "Mynil"
}

var Nil *Null

type handlerFunc func([]string) (string,error)

var handlerTable = map[string]handlerFunc{}

func def(words []string) (string,error) {
	return "",nil
}

func init() {
	handlerTable["def"] = def
}

func repl(treeCh chan Value) {
	for {
		tree := <- treeCh
		fmt.Printf("Tree: <%+v>\n", tree)
	}
}


func treeizeHelper(inCh chan string, curr Value) Value {
	next := <- inCh
	val,err := strconv.ParseInt(next, 10, 64)
	if err == nil {
		// It's an int!
		return &Int{Number: val}
	}

	return Nil
}

func treeize(inCh chan string, outCh chan Value) Value {
	for {
		outCh <- treeizeHelper(inCh, Nil)
	}
}

func tokenize(rawReader io.Reader, outCh chan string) error {
	bufReader := bufio.NewReader(rawReader)
	currTok := make([]byte, 0)
	tokOut := func() {
		trimTok := strings.Trim(string(currTok), " ")
		if len(trimTok) > 0 {
			outCh <- trimTok
		}
		currTok = make([]byte, 0)
	}
	inStr := false
	for {
		b, err := bufReader.ReadByte()
		if err != nil {
			return err
		}
		switch b {
		case '"':
			if inStr {
				// We just ended a string, and therefore a token
				currTok = append(currTok, b)
				tokOut()
				inStr = false
			} else {
				// We're just entering a string
				tokOut()
				currTok = make([]byte,1)
				currTok[0] = '"'
				inStr = true
			}
		case ' ':
			if inStr {
				currTok = append(currTok, b)
			} else {
				tokOut()
			}
		case '\r':
			// We don't talk about this character.
		case '\n':
			tokOut()
		case '(':
			tokOut()
			outCh <- "("
		case ')':
			tokOut()
			outCh <- ")"
		default:
			currTok = append(currTok, b)
		}

	}
}

func main() {
	tokenCh := make(chan string, 256)
	treeCh := make(chan Value, 256)
	go tokenize(os.Stdin, tokenCh)
	go treeize(tokenCh, treeCh)
	repl(treeCh)
}
