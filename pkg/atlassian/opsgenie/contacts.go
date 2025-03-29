package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/contact"
)

func (c *Client) Contacts(user string) []contact.Contact {
	result, e := c.userClient.Contact.List(
		c.context,
		&contact.ListRequest{UserIdentifier: user},
	)
	errors.PanicOnError(e)

	return result.Contact
}
