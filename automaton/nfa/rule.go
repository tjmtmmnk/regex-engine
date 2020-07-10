package nfa

import (
	mapset "github.com/8ayac/golang-set"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
)

type RuleArgs struct {
	From common.State
	C    rune
}

type RuleMap map[RuleArgs]mapset.Set

func NewRuleArgs(from common.State, c rune) RuleArgs {
	return RuleArgs{
		From: from,
		C:    c,
	}
}
