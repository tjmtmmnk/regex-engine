package node

import (
	"fmt"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
	"github.com/tjmtmmnk/regex-engine/automaton/nfa"
)

type Concat struct {
	Ope1 Node
	Ope2 Node
}

func NewConcat(ope1 Node, ope2 Node) *Concat {
	return &Concat{
		Ope1: ope1,
		Ope2: ope2,
	}
}

func (c *Concat) Assemble(ctx *common.Context) *nfa.Fragment {
	fragment := nfa.NewFragment()

	frg1 := c.Ope1.Assemble(ctx)
	frg2 := c.Ope2.Assemble(ctx)

	fragment = frg1.MergeRule(frg2)

	// From frg1's accept state to frg2 ε transition
	for q := range frg1.Accepts.Iter() {
		fragment.AddRule(q.(common.State), 'ε', frg2.Start)
	}

	fragment.Start = frg1.Start
	fragment.Accepts = fragment.Accepts.Union(frg2.Accepts)

	return fragment
}

func (c *Concat) SubtreeString() string {
	return fmt.Sprintf("\x1b[31mconcat(%s, %s\x1b[31m)\x1b[0m", c.Ope1.SubtreeString(), c.Ope2.SubtreeString())
}
