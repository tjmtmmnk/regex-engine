package node

import (
	"fmt"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
	"github.com/tjmtmmnk/regex-engine/automaton/nfa"
)

type Star struct {
	Ope Node
}

func NewStar(ope Node) *Star {
	return &Star{
		Ope: ope,
	}
}

func (s *Star) Assemble(ctx *common.Context) *nfa.Fragment {
	frg := s.Ope.Assemble(ctx)

	fragment := frg.CreateSkeleton()

	for q := range frg.Accepts.Iter() {
		fragment.AddRule(q.(common.State), 'ε', frg.Start)
	}

	q := common.NewState(ctx)

	fragment.AddRule(q, 'ε', frg.Start)

	fragment.Start = q

	fragment.Accepts = fragment.Accepts.Union(frg.Accepts)
	fragment.Accepts.Add(q)

	return fragment
}

func (s *Star) SubtreeString() string {
	return fmt.Sprintf("\x1b[33mstar(%s\x1b[33m)\x1b[0m", s.Ope.SubtreeString())
}
