package regexp

import (
	"github.com/tjmtmmnk/regex-engine/automaton/common"
	"github.com/tjmtmmnk/regex-engine/automaton/dfa"
	"github.com/tjmtmmnk/regex-engine/parser"
)

type Regexp struct {
	s string
	d *dfa.DFA
}

func NewRegexp(s string) *Regexp {
	p := parser.NewParser(s)
	ast := p.GetAST()
	frg := ast.Assemble(common.NewContext())
	nfa := frg.Build()
	dfa := nfa.ToDFA()
	return &Regexp{
		s: s,
		d: dfa,
	}
}

func (r *Regexp) Match(s string) bool {
	runtime := r.d.GetRuntime()
	return runtime.Matching(s)
}
