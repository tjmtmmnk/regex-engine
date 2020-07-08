package node

import "github.com/tjmtmmnk/regex-engine/nfa"

type Plus struct {
	Ope Node
}

func NewPlus(ope Node) *Plus {
	return &Plus{
		Ope: ope,
	}
}

func (c *Plus) Assemble(ctx *nfa.Context) *nfa.Fragment {
	frg := c.Assemble(ctx)

	fragment := frg.CreateSkeleton(ctx)

	for q := range frg.Accepts.Iter() {
		fragment.AddRule(q.(nfa.State), 'ε', frg.Start)
	}

	s := nfa.NewState(ctx)
	fragment.AddRule(s, 'ε', frg.Start)

	fragment.Start = s
	fragment.Accepts.Add(s)

	return fragment
}
