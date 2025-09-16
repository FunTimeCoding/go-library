package jira

func (c *Client) DefaultProjectName() string {
	if c.defaultProjectName == "" {
		panic("default project name not set")
	}

	return c.defaultProjectName
}
