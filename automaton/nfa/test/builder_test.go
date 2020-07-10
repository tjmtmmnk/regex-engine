package test

import (
	"fmt"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
	"github.com/tjmtmmnk/regex-engine/parser"
	"testing"
)

func TestBuild(t *testing.T) {
	p := parser.NewParser("a*")
	ast := p.GetAST()
	fmt.Println(ast.SubtreeString())
	frg := ast.Assemble(common.NewContext())
	nfa := frg.Build()
	fmt.Println(nfa.Start)
	fmt.Println(nfa.Accepts)
	fmt.Println(nfa.Rules)
}
