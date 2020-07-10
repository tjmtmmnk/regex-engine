package common

type Context struct {
	N int
}

func NewContext() *Context {
	return &Context{
		N: -1,
	}
}

func (c *Context) Increment() int {
	c.N++
	return c.N
}
