package node

import (
	"github.com/tjmtmmnk/regex-engine/nfa"
)

type Star struct {
	Ope Node
}

func NewStar(ope Node) *Star {
	return &Star{
		Ope: ope,
	}
}

func (c *Star) Assemble(ctx *nfa.Context) *nfa.Fragment {
	frg := c.Assemble(ctx)

	fragment := frg.CreateSkeleton(ctx)

	for q := range frg.Accepts.Iter() {
		fragment.AddRule(q.(nfa.State), 'ε', frg.Start)
	}

	s := nfa.NewState(ctx)
	fragment.AddRule(s, 'ε', frg.Start)

	fragment.Start = s

	return fragment
}
