package ansible

import "context"

type Client struct {
	context   context.Context
	inventory string
}
