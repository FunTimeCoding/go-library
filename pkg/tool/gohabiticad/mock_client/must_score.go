package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/habitica/response"
)

func (c *Client) MustScore(
	taskID string,
	direction string,
) response.ScoreResult {
	result, e := c.Score(taskID, direction)
	errors.PanicOnError(e)

	return result
}
