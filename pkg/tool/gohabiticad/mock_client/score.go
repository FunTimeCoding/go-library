package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) Score(
	_ string,
	_ string,
) (response.ScoreResult, error) {
	return response.ScoreResult{
		HP:    c.stats.HP,
		MP:    c.stats.MP,
		XP:    c.stats.XP + 10,
		GP:    c.stats.GP + 1,
		Level: c.stats.Level,
	}, nil
}
