package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/habitica/request"
	"github.com/funtimecoding/go-library/pkg/habitica/response"
	"github.com/google/uuid"
)

func (c *Client) CreateTask(
	body request.CreateTaskBody,
) (response.Task, error) {
	task := response.Task{
		ID:   uuid.New().String(),
		Text: body.Text,
		Type: body.Type,
	}
	c.tasks = append(c.tasks, task)

	return task, nil
}
