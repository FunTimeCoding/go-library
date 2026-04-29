package habitica

import "github.com/funtimecoding/go-library/pkg/strings/join"

func (c *Client) Tasks(taskType string) []Task {
	path := "/tasks/user"

	if taskType != "" {
		path = join.Empty(path, "?type=", taskType)
	}

	var result []Task
	c.get(path, &result)

	return result
}
