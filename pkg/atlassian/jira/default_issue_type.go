package jira

func (c *Client) DefaultIssueType() string {
	if c.defaultIssueType == "" {
		panic("default issue type not set")
	}

	return c.defaultIssueType
}
