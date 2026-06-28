package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CloneMachine(
	identifier int64,
	node *string,
	body client.CloneMachineJSONRequestBody,
) string {
	result, e := c.client.CloneMachineWithResponse(
		c.context,
		identifier,
		&client.CloneMachineParams{
			Instance: &c.instance,
			Node:     node,
		},
		body,
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
