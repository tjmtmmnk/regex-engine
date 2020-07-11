package dfa

import "github.com/tjmtmmnk/regex-engine/automaton/common"

type Runtime struct {
	d            *DFA
	currentState common.State
}

func NewRuntime(d *DFA) *Runtime {
	return &Runtime{
		d:            d,
		currentState: d.Start,
	}
}

func (r *Runtime) Matching(str string) bool {
	r.currentState = r.d.Start
	for _, c := range []rune(str) {
		if !r.transit(c) {
			return false
		}
	}
	return r.isAccept()
}

func (r *Runtime) transit(c rune) bool {
	key := NewRuleArgs(r.currentState, c)
	if _, ok := r.d.Rules[key]; ok {
		r.currentState = r.d.Rules[key]
		return true
	}
	return false
}

func (r *Runtime) isAccept() bool {
	return r.d.Accepts.Contains(r.currentState)
}
