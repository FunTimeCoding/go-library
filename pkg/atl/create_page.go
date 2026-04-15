package atl

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreatePage(
	spaceIdentifier string,
	parentIdentifier string,
	title string,
	body string,
) string {
	result, e := c.client.CreatePage(
		c.context,
		client.CreatePageJSONRequestBody{
			SpaceIdentifier:  spaceIdentifier,
			ParentIdentifier: parentIdentifier,
			Title:            title,
			Body:             body,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
