package nfa

import (
	"fmt"
	"github.com/tjmtmmnk/regex-engine/parser"
	"testing"
)

func TestBuild(t *testing.T) {
	p := parser.NewParser("(ab)|c*")
	ast := p.GetAST()
	frg := ast.Assemble(NewContext())
	nfa := frg.Build()
	//fmt.Println(frg)
	fmt.Println(nfa)
}
