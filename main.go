package main

import (
	"io"
	"os"
	"fmt"
	"bufio"
	"strings"
)

type Symbol string

type Value interface{}


type Interpreter struct {
	Globals map[Symbol]Value
	Reader *bufio.Reader
}


type CommandStatus int
const (
	Continue = iota
	Quit
)


type handlerFunc func([]string) (Value,error)

var handlerTable = map[string]handlerFunc{}

func def(words []string) (Value,error) {
	return "",nil	
}

func init() {
	handlerTable["def"] = def
}

// Return the next complete parseable chunk.
func read(interp *Interpreter) (string, CommandStatus) {
	fmt.Printf("$ ")
	val, _ := interp.Reader.ReadString('\n')
	chunk := strings.TrimRight(val, "\r\n")
	return chunk,Continue
}

func eval(interp *Interpreter, cmd string) (Value, CommandStatus) {
	words := strings.Split(cmd, " ")
	if len(words) == 0 {
		return "", Continue
	} else {
		return fmt.Sprintf("Yeah. %+v", words), Continue
	}
}

func print(interp *Interpreter, value Value) {
	fmt.Printf("%+v\n", value)
}

func repl(tokenCh chan string) {
	for {
		token := <- tokenCh
		fmt.Printf("Token: <%v>\n", token)
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
	go tokenize(os.Stdin, tokenCh)
	repl(tokenCh)
}
