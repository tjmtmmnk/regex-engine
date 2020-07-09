package nfa

import mapset "github.com/8ayac/golang-set"

type NFA struct {
	Start   State
	Accepts mapset.Set
	Rules   RuleMap
}

func NewNFA(start State, accepts mapset.Set, rules RuleMap) *NFA {
	return &NFA{
		Start:   start,
		Accepts: accepts,
		Rules:   rules,
	}
}
