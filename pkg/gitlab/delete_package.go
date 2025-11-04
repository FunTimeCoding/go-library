package gitlab

func (c *Client) DeletePackage(
	project int,
	repository int,
) {
	r, e := c.client.Packages.DeleteProjectPackage(project, repository)
	panicOnError(r, e)
}
