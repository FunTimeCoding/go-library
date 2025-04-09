package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/alert"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	rawAlert "github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) Search(query string) []*alert.Alert {
	var result []*alert.Alert
	var start int
	p := c.AlertOption()

	for {
		page, e := c.userClient.Alert.List(
			c.context,
			&rawAlert.ListAlertRequest{
				Limit:  constant.PageLimit,
				Offset: start,
				Query:  query,
			},
		)
		errors.PanicOnError(e)
		result = append(
			result,
			alert.NewSlice(page.Alerts, p, true)...,
		)

		if len(page.Alerts) < constant.PageLimit {
			break
		}

		start += constant.PageLimit
	}

	return c.processor().Process(result)
}
