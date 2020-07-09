package node

import (
	"fmt"
	"github.com/tjmtmmnk/regex-engine/nfa"
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

func (c *Concat) Assemble(ctx *nfa.Context) *nfa.Fragment {
	fragment := nfa.NewFragment(ctx)

	frg1 := c.Ope1.Assemble(ctx)
	frg2 := c.Ope2.Assemble(ctx)

	fragment = frg1.MergeRule(ctx, frg2)

	// From frg1's accept state to frg2 ε transition
	for q := range frg1.Accepts.Iter() {
		fragment.AddRule(q.(nfa.State), 'ε', frg2.Start)
	}

	return fragment
}

func (c *Concat) SubtreeString() string {
	return fmt.Sprintf("\x1b[31mconcat(%s, %s\x1b[31m)\x1b[0m", c.Ope1.SubtreeString(), c.Ope2.SubtreeString())
}
