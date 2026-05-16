package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/habitica/request"
	"github.com/funtimecoding/go-library/pkg/habitica/task"
	"github.com/google/uuid"
)

func (c *Client) CreateTask(b *request.CreateTaskBody) (*task.Task, error) {
	t := task.Stub()
	t.ID = uuid.New().String()
	t.Text = b.Text
	t.Type = b.Type
	c.tasks = append(c.tasks, t)

	return t, nil
}
