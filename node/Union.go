package node

import (
	"fmt"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
	"github.com/tjmtmmnk/regex-engine/automaton/nfa"
)

type Union struct {
	Ope1 Node
	Ope2 Node
}

func NewUnion(ope1 Node, ope2 Node) *Union {
	return &Union{
		Ope1: ope1,
		Ope2: ope2,
	}
}

func (u *Union) Assemble(ctx *common.Context) *nfa.Fragment {
	fragment := nfa.NewFragment(ctx)

	frg1 := u.Ope1.Assemble(ctx)
	frg2 := u.Ope2.Assemble(ctx)

	q := common.NewState(ctx)

	fragment = frg1.MergeRule(ctx, frg2)

	fragment.AddRule(q, 'ε', frg1.Start)
	fragment.AddRule(q, 'ε', frg2.Start)

	// set new start and accepts
	fragment.Start = q
	fragment.Accepts = fragment.Accepts.Union(frg1.Accepts)
	fragment.Accepts = fragment.Accepts.Union(frg2.Accepts)

	return fragment
}

func (u *Union) SubtreeString() string {
	return fmt.Sprintf("\x1b[36munion(%s, %s\x1b[36m)\x1b[0m", u.Ope1.SubtreeString(), u.Ope2.SubtreeString())
}
