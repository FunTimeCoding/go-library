package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/habitica/response"
)

func (c *Client) MustTasks(taskType string) []response.Task {
	result, e := c.Tasks(taskType)
	errors.PanicOnError(e)

	return result
}
