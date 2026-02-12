package proxmox

func WithInsecure() Option {
	return func(c *Client) {
		c.selfSigned = true
	}
}
