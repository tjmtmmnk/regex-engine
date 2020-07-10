package common

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
