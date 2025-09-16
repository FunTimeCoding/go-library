package gitlab

func WithProjects(v []int) Option {
	return func(c *Client) {
		c.projects = v
	}
}
