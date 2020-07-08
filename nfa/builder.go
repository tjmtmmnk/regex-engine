package nfa

import (
	mapset "github.com/8ayac/golang-set"
)

type Fragment struct {
	Start   State
	Accepts mapset.Set
	Rules   RuleMap
}

func NewFragment(ctx *Context) *Fragment {
	return &Fragment{
		Start:   NewState(ctx),
		Accepts: mapset.NewSet(),
		Rules:   RuleMap{},
	}
}

func (f *Fragment) Build()

func (f *Fragment) CreateSkeleton(ctx *Context) (Skeleton *Fragment) {
	Skeleton = NewFragment(ctx)
	Skeleton.Rules = f.Rules
	return
}

func (f *Fragment) AddRule(from State, c rune, to State) {
	rules, ok := f.Rules[NewRuleArgs(from, c)]
	if ok {
		rules.Add(to)
	} else {
		rules = mapset.NewSet(to)
	}
}

func (f *Fragment) MergeRule(ctx *Context, frg *Fragment) (mergedFragment *Fragment) {
	mergedFragment = f.CreateSkeleton(ctx)
	for k, v := range frg.Rules {
		rules, ok := mergedFragment.Rules[k]
		if !ok {
			rules = mapset.NewSet()
		}
		rules = rules.Union(v)
	}
	return
}
