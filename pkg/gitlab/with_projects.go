package gitlab

func WithProjects(v []int64) Option {
	return func(c *Client) {
		c.projects = v
	}
}
