package dfa

import (
	mapset "github.com/8ayac/golang-set"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
)

type DFA struct {
	Start   common.State
	Accepts mapset.Set
	Rules   RuleMap
}

func NewDFA(start common.State, accepts mapset.Set, rules RuleMap) *DFA {
	return &DFA{
		Start:   start,
		Accepts: accepts,
		Rules:   rules,
	}
}
