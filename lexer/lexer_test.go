package lexer

import (
	"github.com/tjmtmmnk/regex-engine/token"
	"testing"
)

func TestScan(t *testing.T) {
	l := NewLexer("(a|+*\a)")
	tokenList := l.Scan()
	expectedList := []token.Type{
		token.LPAREN, token.CHARACTER, token.UNION, token.PLUS, token.STAR, token.CHARACTER, token.RPAREN,
	}
	for i, token := range tokenList {
		if token.Ty != expectedList[i] {
			t.Fail()
		}
	}
}
