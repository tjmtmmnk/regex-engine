package node

import (
	"github.com/tjmtmmnk/regex-engine/automaton/common"
	"github.com/tjmtmmnk/regex-engine/automaton/nfa"
)

type Node interface {
	SubtreeString() string
	Assemble(c *common.Context) *nfa.Fragment
}
