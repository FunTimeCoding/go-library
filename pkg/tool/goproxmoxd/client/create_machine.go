package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateMachine(
	body client.CreateMachineJSONRequestBody,
) string {
	result, e := c.client.CreateMachineWithResponse(
		c.context,
		&client.CreateMachineParams{
			Instance: &c.instance,
		},
		body,
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
