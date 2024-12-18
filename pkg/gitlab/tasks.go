package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/task"
	"gitlab.com/gitlab-org/api/client-go"
	"k8s.io/utils/ptr"
)

func (c *Client) Tasks() []*task.Task {
	t, _, e := c.client.Todos.ListTodos(
		&gitlab.ListTodosOptions{
			State:       ptr.To[string](constant.PendingState),
			ListOptions: constant.DefaultListOptions,
		},
		nil,
	)
	errors.PanicOnError(e)

	return task.NewSlice(t)
}
