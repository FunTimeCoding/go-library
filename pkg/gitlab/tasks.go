package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/task"
	"github.com/xanzy/go-gitlab"
	"k8s.io/utils/ptr"
)

func (c *Client) Tasks() []*task.Task {
	t, _, e := c.client.Todos.ListTodos(
		&gitlab.ListTodosOptions{
			State:       ptr.To[string](PendingState),
			ListOptions: DefaultListOptions,
		},
		nil,
	)
	errors.PanicOnError(e)

	return task.NewSlice(t)
}
