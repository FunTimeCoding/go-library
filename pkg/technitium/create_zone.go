package technitium

import (
	"fmt"
	"net/url"
)

func (c *Client) CreateZone(
	name string,
	zoneType string,
) (string, error) {
	var result createZoneResponse

	return result.Domain, c.get(
		fmt.Sprintf(
			"/zones/create?zone=%s&type=%s",
			url.QueryEscape(name),
			url.QueryEscape(zoneType),
		),
		&result,
	)
}
