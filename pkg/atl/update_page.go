package atl

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) UpdatePage(
	identifier string,
	title string,
	body string,
	message *string,
) string {
	result, e := c.client.UpdatePage(
		c.context,
		identifier,
		client.UpdatePageJSONRequestBody{
			Title:   title,
			Body:    body,
			Message: message,
		},
	)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
