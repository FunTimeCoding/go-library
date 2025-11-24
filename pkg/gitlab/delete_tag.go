package gitlab

func (c *Client) DeleteTag(
	project int64,
	name string,
) {
	r, e := c.client.Tags.DeleteTag(project, name)
	panicOnError(r, e)
}
