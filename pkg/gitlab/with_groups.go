package gitlab

func WithGroups(v []int) Option {
	return func(c *Client) {
		c.groups = append(c.groups, v...)
	}
}
