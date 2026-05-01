package habitica

import (
	"github.com/funtimecoding/go-library/pkg/habitica/response"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (c *Client) Tasks(taskType string) ([]response.Task, error) {
	path := "/tasks/user"

	if taskType != "" {
		path = join.Empty(path, "?type=", taskType)
	}

	var result []response.Task

	return result, c.get(path, &result)
}
