package proxmox

func WithPort(p int) Option {
	return func(c *Client) {
		c.port = p
	}
}
