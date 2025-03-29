package opsgenie

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert"
)

func (c *Client) Open() []*alert.Alert {
	return c.Search(fmt.Sprintf("status:%s", alert.OpenStatus))
}
