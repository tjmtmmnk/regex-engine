package test

import (
	"fmt"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
	"github.com/tjmtmmnk/regex-engine/parser"
	"testing"
)

func TestToDFA(t *testing.T) {
	p := parser.NewParser("a|b")
	ast := p.GetAST()
	frg := ast.Assemble(common.NewContext())
	fmt.Println(ast.SubtreeString())
	nfa := frg.Build()
	//_ = nfa.ToDFA()
	dfa := nfa.ToDFA()
	fmt.Println(dfa)
}
