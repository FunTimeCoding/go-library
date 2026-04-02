package sentry

import "fmt"

func (c *Client) DeleteIssue(
	organization string,
	identifier string,
) {
	c.basic.DeleteBytes(
		fmt.Sprintf(
			"organizations/%s/issues/%s",
			organization,
			identifier,
		),
	)
}
