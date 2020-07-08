package nfa

type State struct {
	N int
}

func NewState(ctx *Context) State {
	return State{
		N: ctx.Increment(),
	}
}
