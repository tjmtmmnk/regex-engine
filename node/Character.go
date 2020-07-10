package node

import (
	"fmt"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
	"github.com/tjmtmmnk/regex-engine/automaton/nfa"
)

type Character struct {
	V rune
}

func NewCharacter(r rune) *Character {
	return &Character{
		V: r,
	}
}

func (c *Character) Assemble(ctx *common.Context) *nfa.Fragment {
	fragment := nfa.NewFragment(ctx)

	q1 := common.NewState(ctx)
	q2 := common.NewState(ctx)

	fragment.AddRule(q1, c.V, q2)

	fragment.Start = q1
	fragment.Accepts.Add(q2)

	return fragment
}

func (c *Character) SubtreeString() string {
	return fmt.Sprintf("\x1b[32mcharacter('%s')\x1b[32m", string(c.V))
}
