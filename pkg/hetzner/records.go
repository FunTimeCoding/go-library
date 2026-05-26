package hetzner

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/hetzner/record"
	"github.com/funtimecoding/go-library/pkg/hetzner/zone"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func (c *Client) Records(z *zone.Zone) []*record.Record {
	result, e := c.client.Zone.AllRRSetsWithOpts(
		c.context,
		z.Raw,
		hcloud.ZoneRRSetListOpts{},
	)
	errors.PanicOnError(e)

	return record.NewSlice(result, z)
}
