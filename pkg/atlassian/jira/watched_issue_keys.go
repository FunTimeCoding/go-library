package jira

func (c *Client) WatchedIssueKeys() []string {
	var result []string

	for _, i := range c.SearchV3(
		"issue IN watchedIssues() ORDER BY key ASC",
	) {
		result = append(result, i.Key)
	}

	return result
}
