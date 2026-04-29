package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/task"
)

func (c *Client) MustTasks() []*task.Task {
	result, e := c.Tasks()
	errors.PanicOnError(e)

	return result
}
