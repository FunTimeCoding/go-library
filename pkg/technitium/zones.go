package technitium

import "github.com/funtimecoding/go-library/pkg/technitium/zone"

func (c *Client) Zones() ([]*zone.Zone, error) {
	var result zonesResponse

	return result.Zones, c.get("/zones/list", &result)
}
