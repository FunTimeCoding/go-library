package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) AddTask(task response.Task) {
	c.tasks = append(c.tasks, task)
}
