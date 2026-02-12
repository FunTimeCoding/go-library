package proxmox

func WithUser(
	user string,
	password string,
) Option {
	return func(c *Client) {
		c.user = user
		c.password = password
	}
}
