package discord

func (c *Client) Agent() string {
	return c.client.UserAgent
}
