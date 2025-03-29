package opsgenie

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert/option"

func (c *Client) AlertOption() *option.Alert {
	return option.New(
		c.TeamMap(),
		c.UserMap(),
		c.webHost,
		c.shortAlert,
		c.shortUser,
		c.descriptionToName,
		c.parseDescription,
	)
}
