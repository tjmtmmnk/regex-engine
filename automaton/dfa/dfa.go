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

func (d *DFA) GetRuntime() *Runtime {
	return NewRuntime(d)
}

func (d *DFA) Minimize() {
	// create DFA all states
	states := mapset.NewSet(d.Start)
	for _, state := range d.Rules {
		states.Add(state)
	}

	seen := map[common.State]common.State{}
	for i := 0; i < states.Cardinality(); i++ {
		toState := common.NewStateWithNumber(i)
		for j := i + 1; j < states.Cardinality(); j++ {
			fromState := common.NewStateWithNumber(j)
			if !d.isEquivalent(fromState, toState) {
				continue
			}
			if _, ok := seen[fromState]; ok {
				continue
			}
			seen[fromState] = toState
			d.mergeState(fromState, toState)
		}
	}
}

func (d *DFA) mergeState(from, to common.State) {
	rules := d.Rules
	for args, state := range rules {
		// delete redundant state and its transition
		if state == from {
			rules[args] = to
		}
		if args.From == from {
			delete(rules, args)
		}
	}
}

func (d *DFA) isEquivalent(q1, q2 common.State) bool {
	// both accept or not
	if !((d.Accepts.Contains(q1) && d.Accepts.Contains(q2)) ||
		(!d.Accepts.Contains(q1) && !d.Accepts.Contains(q2))) {
		return false
	}
	for args := range d.Rules {
		if args.From != q1 {
			continue
		}
		// same transition by same symbol
		if d.Rules[NewRuleArgs(q1, args.C)] != d.Rules[NewRuleArgs(q2, args.C)] {
			return false
		}
	}
	return true
}
