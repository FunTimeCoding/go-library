package gitlab

func WithGroups(v []int64) Option {
	return func(c *Client) {
		c.groups = append(c.groups, v...)
	}
}
