package gitlab

func WithGroups(v []int) OptionFunc {
	return func(c *Client) {
		c.groups = append(c.groups, v...)
	}
}
