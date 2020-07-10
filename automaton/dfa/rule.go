package dfa

import (
	"github.com/tjmtmmnk/regex-engine/automaton/common"
)

type RuleMap map[RuleArgs]common.State

type RuleArgs struct {
	From common.State
	C    rune
}

func NewRuleArgs(from common.State, in rune) RuleArgs {
	return RuleArgs{
		From: from,
		C:    in,
	}
}
