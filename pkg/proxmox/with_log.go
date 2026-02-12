package proxmox

func WithLog() Option {
	return func(c *Client) {
		c.log = true
	}
}
