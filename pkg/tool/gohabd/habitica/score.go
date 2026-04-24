package habitica

func (c *Client) Score(
	taskID string,
	direction string,
) ScoreResult {
	if direction == "" {
		direction = "up"
	}

	var result ScoreResult
	c.post("/tasks/"+taskID+"/score/"+direction, nil, &result)

	return result
}
