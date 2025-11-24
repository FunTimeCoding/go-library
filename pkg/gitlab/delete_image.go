package gitlab

func (c *Client) DeleteImage(
	project int64,
	repository int64,
	tag string,
) {
	r, e := c.client.ContainerRegistry.DeleteRegistryRepositoryTag(
		project,
		repository,
		tag,
	)
	panicOnError(r, e)
}
