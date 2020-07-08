package nfa

import (
	mapset "github.com/8ayac/golang-set"
)

type RuleArgs struct {
	From State
	C    rune
}

type RuleMap map[RuleArgs]mapset.Set

func NewRuleArgs(from State, c rune) RuleArgs {
	return RuleArgs{
		From: from,
		C:    c,
	}
}
