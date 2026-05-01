package habitica

import (
	"github.com/funtimecoding/go-library/pkg/habitica/request"
	"github.com/funtimecoding/go-library/pkg/habitica/response"
)

func (c *Client) CreateTask(body request.CreateTaskBody) (response.Task, error) {
	var result response.Task

	return result, c.post("/tasks/user", body, &result)
}
