package opsgenie

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert"

func (c *Client) Open() []*alert.Alert {
	return c.Search("status:%s", alert.OpenStatus)
}
