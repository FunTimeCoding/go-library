package opsgenie

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert"

func (c *Client) OpenUnowned() []*alert.Alert {
	return c.Search("status:%s owner:NULL", alert.OpenStatus)
}
