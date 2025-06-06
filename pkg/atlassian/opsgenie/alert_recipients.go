package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

func (c *Client) AlertRecipients(identifier string) []alert.AlertRecipient {
	result, e := c.userClient.Alert.ListAlertRecipients(
		c.context,
		&alert.ListAlertRecipientRequest{
			IdentifierType:  alert.ALERTID,
			IdentifierValue: identifier,
		},
	)
	errors.PanicOnError(e)

	return result.AlertRecipients
}
