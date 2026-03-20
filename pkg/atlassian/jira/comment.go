package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Comment(
	key string,
	body string,
) {
	_, r, e := c.client.Issue.AddCommentWithContext(
		c.context,
		key,
		&jira.Comment{Body: body},
	)
	panicOnError(r, e)
}
