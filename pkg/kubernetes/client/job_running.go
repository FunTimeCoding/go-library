package client

func (c *Client) JobRunning(
	namespace string,
	name string,
) bool {
	j := c.Job(namespace, name)

	if j == nil {
		return false
	}

	return j.Status.CompletionTime == nil && j.Status.Failed == 0
}
