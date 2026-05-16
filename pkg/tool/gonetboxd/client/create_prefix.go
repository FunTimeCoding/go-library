package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreatePrefix(
	prefix string,
	site string,
	description string,
) string {
	body := client.CreatePrefixRequest{Prefix: prefix}

	if site != "" {
		body.Site = &site
	}

	if description != "" {
		body.Description = &description
	}

	result, e := c.client.CreatePrefix(c.context, body)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
