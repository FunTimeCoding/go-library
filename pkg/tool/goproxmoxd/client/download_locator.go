package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) DownloadLocator(
	node string,
	storage string,
	body client.DownloadLocatorJSONRequestBody,
) string {
	result, e := c.client.DownloadLocatorWithResponse(
		c.context,
		node,
		storage,
		&client.DownloadLocatorParams{
			Instance: &c.instance,
		},
		body,
	)
	errors.PanicOnError(e)

	return web.ReadString(result.HTTPResponse)
}
