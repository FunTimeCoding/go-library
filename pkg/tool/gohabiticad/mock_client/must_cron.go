package mock_client

import "github.com/funtimecoding/go-library/pkg/habitica/cron"

func (c *Client) MustCron() *cron.Cron {
	return cron.New()
}
