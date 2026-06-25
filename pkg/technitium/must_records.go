package technitium

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/technitium/record"
)

func (c *Client) MustRecords(
	domain string,
	listZone bool,
) []*record.Record {
	result, e := c.Records(domain, listZone)
	errors.PanicOnError(e)

	return result
}
