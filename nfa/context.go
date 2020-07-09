package nfa

import "fmt"

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
	fmt.Println(c.N)
	return c.N
}
