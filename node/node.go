package node

import (
	"github.com/tjmtmmnk/regex-engine/nfa"
)

type Node interface {
	SubtreeString() string
	Assemble(c *nfa.Context) *nfa.Fragment
}
