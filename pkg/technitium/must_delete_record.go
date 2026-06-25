package technitium

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDeleteRecord(
	domain string,
	recordType string,
	value string,
) {
	errors.PanicOnError(c.DeleteRecord(domain, recordType, value))
}
