package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) Cron() (response.CronResult, error) {
	return response.CronResult{RolledOver: false}, nil
}
