package alertmanager

func (c *Client) SimpleSilence(alert string) {
	c.CreateSilence(alert, "", 0)
}
