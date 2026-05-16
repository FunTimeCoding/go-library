package habitica

import (
	"github.com/funtimecoding/go-library/pkg/habitica/task"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (c *Client) Tasks(taskType string) ([]*task.Task, error) {
	path := "/tasks/user"

	if taskType != "" {
		path = join.Empty(path, "?type=", taskType)
	}

	var result []*task.Task

	return result, c.get(path, &result)
}
