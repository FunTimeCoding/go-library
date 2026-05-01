package habitica

import "github.com/funtimecoding/go-library/pkg/habitica/response"

func (c *Client) Cron() (response.CronResult, error) {
	before, e := c.user()

	if e != nil {
		return response.CronResult{}, e
	}

	if !before.NeedsCron {
		return response.CronResult{
			LastCron: before.LastCron,
		}, nil
	}

	if f := c.postDiscard("/cron"); f != nil {
		return response.CronResult{}, f
	}

	after, g := c.user()

	if g != nil {
		return response.CronResult{}, g
	}

	return response.CronResult{
		RolledOver: true,
		LastCron:   after.LastCron,
		Before:     &before.Stats,
		After:      &after.Stats,
	}, nil
}
