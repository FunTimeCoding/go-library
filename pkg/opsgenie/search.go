package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/opsgenie/alert"
	"github.com/funtimecoding/go-library/pkg/opsgenie/constant"
	raw "github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) Search(query string) []*alert.Alert {
	var result []*alert.Alert
	var start int
	p := c.AlertOption()

	for {
		page, e := c.userClient.Alert.List(
			c.context,
			&raw.ListAlertRequest{
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

	return result
}
