package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) AlertLogs(identifier string) []alert.AlertLog {
	result, e := c.userClient.Alert.ListAlertLogs(
		c.context,
		&alert.ListAlertLogsRequest{
			IdentifierType:  alert.ALERTID,
			IdentifierValue: identifier,
		},
	)
	errors.PanicOnError(e)

	return result.AlertLog
}
