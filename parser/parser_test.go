package parser

import (
	"fmt"
	"github.com/tjmtmmnk/regex-engine/token"
	"testing"
)

func TestMoveWithValidation(t *testing.T) {
	p := NewParser("a+")
	if p.look.V != 'a' || p.look.Ty != token.CHARACTER {
		t.Fail()
	}
	p.moveWithValidation(token.CHARACTER)
	if p.look.V != '+' || p.look.Ty != token.PLUS {
		t.Fail()
	}
}

func TestGetAST(t *testing.T) {
	p := NewParser("(ab+)|c*")
	ast := p.GetAST()
	fmt.Println(ast.SubtreeString())
}
