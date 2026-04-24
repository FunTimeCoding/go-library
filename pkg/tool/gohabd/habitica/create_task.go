package habitica

func (c *Client) CreateTask(body CreateTaskBody) Task {
	var result Task
	c.post("/tasks/user", body, &result)

	return result
}
