package node

import (
	"fmt"
	"github.com/tjmtmmnk/regex-engine/nfa"
)

type Plus struct {
	Ope Node
}

func NewPlus(ope Node) *Plus {
	return &Plus{
		Ope: ope,
	}
}

func (p *Plus) Assemble(ctx *nfa.Context) *nfa.Fragment {
	frg := p.Ope.Assemble(ctx)

	fragment := frg.CreateSkeleton(ctx)

	for q := range frg.Accepts.Iter() {
		fragment.AddRule(q.(nfa.State), 'ε', frg.Start)
	}

	s := nfa.NewState(ctx)
	fragment.AddRule(s, 'ε', frg.Start)

	fragment.Start = s
	
	fragment.Accepts.Union(frg.Accepts)
	fragment.Accepts.Add(s)

	return fragment
}

func (p *Plus) SubtreeString() string {
	return fmt.Sprintf("\x1b[33mplus(%s\x1b[33m)\x1b[0m", p.Ope.SubtreeString())
}
