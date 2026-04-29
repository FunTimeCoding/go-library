package habitica

func (c *Client) Tags() []Tag {
	var result []Tag
	c.get("/tags", &result)

	return result
}
