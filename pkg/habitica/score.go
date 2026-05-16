package habitica

import (
	"github.com/funtimecoding/go-library/pkg/habitica/score"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (c *Client) Score(
	taskID string,
	direction string,
) (*score.Score, error) {
	if direction == "" {
		direction = "up"
	}

	var result *score.Score

	return result, c.post(
		join.Empty("/tasks/", taskID, "/score/", direction),
		nil,
		&result,
	)
}
