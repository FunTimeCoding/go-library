package flag

import "github.com/funtimecoding/go-library/pkg/monitor/gorilla/router/client"

func New(
	identifier string,
	c *client.Client,
) *Flag {
	return &Flag{Identifier: identifier, By: []*client.Client{c}}
}
