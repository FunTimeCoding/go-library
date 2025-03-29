package opsgenie

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert"
)

func (c *Client) OpenUnowned() []*alert.Alert {
	return c.Search(
		fmt.Sprintf("status:%s owner:NULL", alert.OpenStatus),
	)
}
