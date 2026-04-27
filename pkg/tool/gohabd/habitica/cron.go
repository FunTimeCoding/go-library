package habitica

func (c *Client) Cron() CronResult {
	before := c.user()

	if !before.NeedsCron {
		return CronResult{
			LastCron: before.LastCron,
		}
	}

	c.postDiscard("/cron")
	after := c.user()

	return CronResult{
		RolledOver: true,
		LastCron:   after.LastCron,
		Before:     &before.Stats,
		After:      &after.Stats,
	}
}
