package client

func (c *Client) Consume() []string {
	result := c.receive
	c.receive = []string{}

	return result
}
