package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) MustCron() response.CronResult {
	return response.CronResult{RolledOver: false}
}
