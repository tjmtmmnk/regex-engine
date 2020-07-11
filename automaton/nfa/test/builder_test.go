package test

import (
	"fmt"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
	"github.com/tjmtmmnk/regex-engine/parser"
	"testing"
)

func TestBuild(t *testing.T) {
	p := parser.NewParser("a|b")
	ast := p.GetAST()
	frg := ast.Assemble(common.NewContext())
	nfa := frg.Build()
	fmt.Println(nfa)
}
