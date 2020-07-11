package test

import (
	"fmt"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
	"github.com/tjmtmmnk/regex-engine/parser"
	"testing"
)

func TestMatching(t *testing.T) {
	p := parser.NewParser("bana(na)*")
	ast := p.GetAST()
	frg := ast.Assemble(common.NewContext())
	fmt.Println(ast.SubtreeString())
	nfa := frg.Build()
	dfa := nfa.ToDFA()
	r := dfa.GetRuntime()

	for _, s := range []string{"bana", "banana", "banana", "banaNa", "apple"} {
		MATCHED := r.Matching(s)
		if MATCHED {
			fmt.Printf("%s\t=> Matched!\n", s)
		} else {
			fmt.Printf("%s\t=> NOT Matched...\n", s)
		}
	}
}
