package nfa

import (
	"fmt"
	mapset "github.com/8ayac/golang-set"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
	"github.com/tjmtmmnk/regex-engine/automaton/dfa"
)

type NFA struct {
	Start   common.State
	Accepts mapset.Set
	Rules   RuleMap
}

func NewNFA(start common.State, accepts mapset.Set, rules RuleMap) *NFA {
	return &NFA{
		Start:   start,
		Accepts: accepts,
		Rules:   rules,
	}
}

func (nfa *NFA) ToWithoutEpsilon() {
	if nfa.Accepts.IsSubset(nfa.epsilonClosure(nfa.Start)) {
		nfa.Accepts.Add(nfa.Start)
	}
	nfa.Rules = nfa.removeEpsilonRule()
}

func (nfa *NFA) ToDFA() *dfa.DFA {
	nfa.ToWithoutEpsilon()
	start, accepts, rules := nfa.constructSubset()
	return dfa.NewDFA(start, accepts, rules)
}

func (nfa *NFA) constructSubset() (dStart common.State, dAccepts mapset.Set, dRules dfa.RuleMap) {
	start := nfa.Start
	accepts := nfa.Accepts
	rules := nfa.Rules

	dStart = common.NewStateWithNumber(0)
	dAccepts = mapset.NewSet()
	dRules = dfa.RuleMap{}
	dStates := dfa.StatesMap{}
	dStates[mapset.NewSet(start)] = common.NewStateWithNumber(0)

	queue := mapset.NewSet(mapset.NewSet(start))

	for queue.Cardinality() > 0 {
		targetStates := queue.Pop().(mapset.Set)

		if accepts.Intersect(targetStates).Cardinality() > 0 {
			dAccepts.Add(dStates.GetState(targetStates))
		}

		for c := range nfa.allSymbols().Iter() {
			dnext := mapset.NewSet()
			for q := range targetStates.Iter() {
				d, ok := rules[NewRuleArgs(q.(common.State), c.(rune))]
				if ok {
					dnext = dnext.Union(d)
				}
			}

			if dnext.Cardinality() == 0 {
				continue
			}

			if !dStates.HaveKey(dnext) {
				queue.Add(dnext)
				dStates[dnext] = common.NewStateWithNumber(len(dStates))
			}

			for k := range dStates {
				if k.Equal(dnext) {
					dnext = k
				}
			}
			dRules[dfa.NewRuleArgs(dStates[targetStates], c.(rune))] = dStates[dnext]
		}
	}
	return
}

func (nfa *NFA) allStates() (states mapset.Set) {
	states = mapset.NewSet()
	for rule := range nfa.Rules {
		states.Add(rule.From)
	}

	return
}

func (nfa *NFA) allSymbols() (symbols mapset.Set) {
	symbols = mapset.NewSet()
	for rule := range nfa.Rules {
		symbols.Add(rule.C)
	}
	return
}

func (nfa *NFA) epsilonClosure(state common.State) (reachableStates mapset.Set) {
	// state's self
	reachableStates = mapset.NewSet(state)

	if states, ok := nfa.getTransitionedStates(state, 'ε'); ok {
		reachableStates = reachableStates.Union(states)
	}

	return
}

// ∀q ∈ Q, ∀c ∈ Σ s.t. δ(q,a) = ε-CL(∪_{q' ∈ ε-CL(q)} δ(q', c))
func (nfa *NFA) removeEpsilonRule() (newRule RuleMap) {
	newRule = RuleMap{}
	states, symbols := nfa.allStates(), nfa.allSymbols()
	symbols.Remove('ε')

	for q := range states.Iter() {
		for c := range symbols.Iter() {
			q := q.(common.State)
			c := c.(rune)
			for ec := range nfa.epsilonClosure(q).Iter() {
				expand := nfa.epsilonExpand(ec.(common.State), c)
				s, ok := newRule[NewRuleArgs(q, c)]
				if !ok {
					s = mapset.NewSet()
				}
				newRule[NewRuleArgs(q, c)] = s.Union(expand)
			}
		}
	}

	// remove empty set
	for k := range newRule {
		if newRule[k].Cardinality() == 0 {
			delete(newRule, k)
		}
	}
	return
}

// ε* -> symbol -> ε*
func (nfa *NFA) epsilonExpand(state common.State, symbol rune) mapset.Set {
	firstEpsilonStates := nfa.epsilonClosure(state)

	symbolStates := mapset.NewSet()
	for q := range firstEpsilonStates.Iter() {
		if states, ok := nfa.getTransitionedStates(q.(common.State), symbol); ok {
			symbolStates = symbolStates.Union(states)
		}
	}

	secondEpsilonStates := mapset.NewSet()
	for q := range symbolStates.Iter() {
		secondEpsilonStates = secondEpsilonStates.Union(nfa.epsilonClosure(q.(common.State)))
	}

	return secondEpsilonStates
}

func (nfa *NFA) getTransitionedStates(state common.State, symbol rune) (mapset.Set, bool) {
	states, ok := nfa.Rules[NewRuleArgs(state, symbol)]
	if ok {
		return states, true
	}
	return nil, false
}
