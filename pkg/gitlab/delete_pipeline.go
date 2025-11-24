package gitlab

func (c *Client) DeletePipeline(
	project int64,
	pipeline int64,
) {
	r, e := c.client.Pipelines.DeletePipeline(project, pipeline)

	if r != nil && r.StatusCode == 404 {
		// Do not panic
		return
	}

	panicOnError(r, e)
}
