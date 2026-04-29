package client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/generated/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) Tasks(taskType string) string {
	var p *client.GetTasksParams

	if taskType != "" {
		t := client.GetTasksParamsType(taskType)
		p = &client.GetTasksParams{Type: &t}
	}

	result, e := c.client.GetTasks(c.context, p)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
