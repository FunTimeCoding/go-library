package example

import "context"

func ContextVar() {
	c := context.Background() // want `variable c of type context.Context should be named x`
	_ = c
}
