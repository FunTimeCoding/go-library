package proxmox

func WithToken(
	token string,
	secret string,
) Option {
	return func(c *Client) {
		c.token = token
		c.secret = secret
	}
}
