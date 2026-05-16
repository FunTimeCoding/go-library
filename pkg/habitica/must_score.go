package habitica

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/habitica/score"
)

func (c *Client) MustScore(
	taskID string,
	direction string,
) *score.Score {
	result, e := c.Score(taskID, direction)
	errors.PanicOnError(e)

	return result
}
