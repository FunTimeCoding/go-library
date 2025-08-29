package power_feed

import "github.com/netbox-community/go-netbox/v4"

type Feed struct {
	Identifier int32
	Name       string

	Raw *netbox.PowerFeed
}
