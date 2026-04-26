package habitica

import "github.com/funtimecoding/go-library/pkg/strings/join"

func (c *Client) Score(
	taskID string,
	direction string,
) ScoreResult {
	if direction == "" {
		direction = "up"
	}

	var result ScoreResult
	c.post(join.Empty("/tasks/", taskID, "/score/", direction), nil, &result)

	return result
}
