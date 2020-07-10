package dfa

import (
	mapset "github.com/8ayac/golang-set"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
)

type StatesMap map[mapset.Set]common.State

func (s StatesMap) GetState(key mapset.Set) common.State {
	if s.HaveKey(key) {
		for k := range s {
			if k.Equal(key) {
				return s[k]
			}
		}
	}
	return common.State{}
}

func (s StatesMap) HaveKey(key mapset.Set) bool {
	// check whether equal
	for k := range s {
		if k.Equal(key) {
			return true
		}
	}
	return false
}
