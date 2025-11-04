package gitlab

func (c *Client) DeletePipeline(
	project int,
	pipeline int,
) {
	r, e := c.client.Pipelines.DeletePipeline(project, pipeline)

	if r != nil && r.StatusCode == 404 {
		// Do not panic
		return
	}

	panicOnError(r, e)
}
