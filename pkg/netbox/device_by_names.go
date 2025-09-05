package netbox

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netbox/device"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (c *Client) DeviceByNames(
	n []string,
) (*device.Device, error) {
	var result *device.Device

	for _, name := range n {
		devices := c.DevicesByName(name)

		if len(devices) == 0 {
			continue
		}

		if len(devices) > 1 {
			var identifiers []string

			for _, d := range devices {
				identifiers = append(
					identifiers,
					fmt.Sprintf("%d", d.Raw.Id),
				)
			}

			return nil, fmt.Errorf(
				"more than one device found with name '%s'. IDs: %s",
				name,
				join.Comma(identifiers),
			)
		}

		result = devices[0]

		break
	}

	if result == nil {
		return nil, fmt.Errorf(
			"no device found matching names: %s",
			join.Comma(n),
		)
	}

	return result, nil
}
