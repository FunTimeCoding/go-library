package habitica

import (
	"github.com/funtimecoding/go-library/pkg/habitica/response"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (c *Client) Score(
	taskID string,
	direction string,
) (response.ScoreResult, error) {
	if direction == "" {
		direction = "up"
	}

	var result response.ScoreResult

	return result, c.post(
		join.Empty("/tasks/", taskID, "/score/", direction),
		nil,
		&result,
	)
}
