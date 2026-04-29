package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) CreateTask(
	taskType string,
	text string,
	notes string,
) string {
	body := client.CreateTaskJSONRequestBody{
		Type: client.CreateTaskRequestType(taskType),
		Text: text,
	}

	if notes != "" {
		body.Notes = &notes
	}

	result, e := c.client.CreateTask(c.context, body)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
