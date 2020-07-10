package nfa

import (
	"fmt"
	mapset "github.com/8ayac/golang-set"
	"github.com/tjmtmmnk/regex-engine/automaton/common"
	"reflect"
)

type RuleArgs struct {
	From common.State
	C    rune
}

type RuleMap map[RuleArgs]mapset.Set

func NewRuleArgs(from common.State, c rune) RuleArgs {
	return RuleArgs{
		From: from,
		C:    c,
	}
}

func (r RuleMap) String() string {
	s := ""

	keys := reflect.ValueOf(r).MapKeys()
	for i, k := range keys {
		from := k.FieldByName("From").Interface().(common.State)
		c := k.FieldByName("C").Interface().(rune)
		dst := r[NewRuleArgs(from, c)]
		s += fmt.Sprintf("%s\t--['%c']-->\t%s", from, c, dst)
		if i+1 < len(keys) {
			s += "\n"
		}
	}
	return s
}
