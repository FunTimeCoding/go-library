package habitica

func (c *Client) Tasks(taskType string) []Task {
	path := "/tasks/user"

	if taskType != "" {
		path += "?type=" + taskType
	}

	var result []Task
	c.get(path, &result)

	return result
}
