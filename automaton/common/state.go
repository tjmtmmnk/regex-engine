package common

import "fmt"

type State struct {
	N int
}

func NewState(ctx *Context) State {
	return State{
		N: ctx.Increment(),
	}
}

func NewStateWithNumber(n int) State {
	return State{
		N: n,
	}
}

func (s State) String() string {
	return fmt.Sprintf("q%d", s.N)
}
