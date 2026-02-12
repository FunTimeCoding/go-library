package proxmox

func WithVerbose() Option {
	return func(c *Client) {
		c.verbose = true
	}
}
