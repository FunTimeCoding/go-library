package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) AlertNotes(identifier string) []alert.AlertNote {
	result, e := c.userClient.Alert.ListAlertNotes(
		c.context,
		&alert.ListAlertNotesRequest{
			IdentifierType:  alert.ALERTID,
			IdentifierValue: identifier,
		},
	)
	errors.PanicOnError(e)

	return result.AlertLog
}
