package gitlab

func WithProjects(v []int) OptionFunc {
	return func(c *Client) {
		c.projects = v
	}
}
