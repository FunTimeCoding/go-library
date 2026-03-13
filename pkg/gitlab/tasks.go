package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/task"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) Tasks() []*task.Task {
	result, r, e := c.client.Todos.ListTodos(
		&gitlab.ListTodosOptions{
			State:       new(constant.PendingState),
			ListOptions: constant.DefaultListOptions,
		},
		nil,
	)
	panicOnError(r, e)

	return task.NewSlice(result)
}
