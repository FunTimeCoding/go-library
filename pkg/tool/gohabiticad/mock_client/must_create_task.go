package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/habitica/request"
	"github.com/funtimecoding/go-library/pkg/habitica/task"
)

func (c *Client) MustCreateTask(b *request.CreateTaskBody) *task.Task {
	result, e := c.CreateTask(b)
	errors.PanicOnError(e)

	return result
}
