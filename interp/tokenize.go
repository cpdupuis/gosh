package interp

import (
	"io"
	"bufio"
	"strings"
)

func Tokenize(rawReader io.Reader, outCh chan string) error {
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
