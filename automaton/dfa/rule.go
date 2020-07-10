package dfa

import (
	"fmt"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
	"reflect"
)

type RuleMap map[RuleArgs]common.State

type RuleArgs struct {
	From common.State
	C    rune
}

func NewRuleArgs(from common.State, in rune) RuleArgs {
	return RuleArgs{
		From: from,
		C:    in,
	}
}

func (r RuleMap) String() string {
	s := ""

	keys := reflect.ValueOf(r).MapKeys()
	for i, k := range keys {
		from := k.FieldByName("From").Interface().(common.State)
		c := k.FieldByName("C").Interface().(rune)
		dst := r[NewRuleArgs(from, c)]
		s += fmt.Sprintf("s%d\t--['%c']-->\t%s", from, c, dst)
		if i+1 < len(keys) {
			s += "\n"
		}
	}
	return s
}
