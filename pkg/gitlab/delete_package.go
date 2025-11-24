package gitlab

func (c *Client) DeletePackage(
	project int64,
	repository int64,
) {
	r, e := c.client.Packages.DeleteProjectPackage(project, repository)
	panicOnError(r, e)
}
