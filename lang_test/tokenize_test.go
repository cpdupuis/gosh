package lang_test
import (
	"github.com/cpdupuis/gosh/lang"
	"strings"
	"testing"
)

func TestSomeTokens(t *testing.T) {
	reader := strings.NewReader("The rain (in Spain)  stays \n((mainly)) \"on the plain.\"")
	ch := make(chan string, 256)
	go lang.Tokenize(reader, ch)
	if <-ch != "The" {
		t.Fail()
	}
	if <-ch != "rain" {
		t.Fail()
	}
	if <-ch != "(" {
		t.Fail()
	}
	if <-ch != "in" {
		t.Fail()
	}
	if <-ch != "Spain" {
		t.Fail()
	}
	if <-ch != ")" {
		t.Fail()
	}
	if <-ch != "stays" {
		t.Fail()
	}
	if <-ch != "(" {
		t.Fail()
	}
	if <-ch != "(" {
		t.Fail()
	}
	if <-ch != "mainly" {
		t.Fail()
	}
	if <-ch != ")" {
		t.Fail()
	}
	if <-ch != ")" {
		t.Fail()
	}
	if <-ch != "\"on the plain.\"" {
		t.Fail()
	}
}