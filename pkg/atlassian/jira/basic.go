package jira

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/basic"

func (c *Client) Basic() *basic.Client {
	return c.basic
}
