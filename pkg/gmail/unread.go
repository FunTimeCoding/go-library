package gmail

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"google.golang.org/api/gmail/v1"
)

func (c *Client) Unread() *gmail.ListMessagesResponse {
	result, e := c.service.Users.Messages.List(
		"me",
	).Q("is:unread").MaxResults(10).Do()
	errors.PanicOnError(e)

	return result
}
