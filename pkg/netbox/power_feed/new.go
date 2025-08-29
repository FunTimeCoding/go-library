package power_feed

import "github.com/netbox-community/go-netbox/v4"

func New(d *netbox.PowerFeed) *Feed {
	return &Feed{Identifier: d.GetId(), Name: d.GetName(), Raw: d}
}
