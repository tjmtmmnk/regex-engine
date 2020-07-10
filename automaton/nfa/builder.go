package nfa

import (
	mapset "github.com/8ayac/golang-set"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
)

type Fragment struct {
	Start   common.State
	Accepts mapset.Set
	Rules   RuleMap
}

func NewFragment(ctx *common.Context) *Fragment {
	return &Fragment{
		Start:   common.NewState(ctx),
		Accepts: mapset.NewSet(),
		Rules:   RuleMap{},
	}
}

func (f *Fragment) Build() *NFA {
	return NewNFA(f.Start, f.Accepts, f.Rules)
}

func (f *Fragment) CreateSkeleton(ctx *common.Context) (Skeleton *Fragment) {
	Skeleton = NewFragment(ctx)
	Skeleton.Rules = f.Rules
	return
}

func (f *Fragment) AddRule(from common.State, c rune, to common.State) {
	_, ok := f.Rules[NewRuleArgs(from, c)]
	if ok {
		f.Rules[NewRuleArgs(from, c)].Add(to)
	} else {
		f.Rules[NewRuleArgs(from, c)] = mapset.NewSet(to)
	}
}

func (f *Fragment) MergeRule(ctx *common.Context, frg *Fragment) (mergedFragment *Fragment) {
	mergedFragment = f.CreateSkeleton(ctx)
	for k, v := range frg.Rules {
		_, ok := mergedFragment.Rules[k]
		if !ok {
			mergedFragment.Rules[k] = mapset.NewSet()
		}
		mergedFragment.Rules[k] = mergedFragment.Rules[k].Union(v)
	}
	return
}
