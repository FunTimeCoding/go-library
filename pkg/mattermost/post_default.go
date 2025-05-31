package mattermost

func (c *Client) PostDefault(text string) {
	if c.channel != nil {
		c.Post(c.channel, text)
	}
}
