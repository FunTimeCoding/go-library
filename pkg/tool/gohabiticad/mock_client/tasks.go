package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) Tasks(taskType string) ([]response.Task, error) {
	if taskType == "" {
		return c.tasks, nil
	}

	var result []response.Task

	for _, t := range c.tasks {
		if t.Type == taskType {
			result = append(result, t)
		}
	}

	return result, nil
}
