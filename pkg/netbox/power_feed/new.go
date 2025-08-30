package power_feed

import "github.com/netbox-community/go-netbox/v4"

func New(v *netbox.PowerFeed) *Feed {
	return &Feed{Identifier: v.GetId(), Name: v.GetName(), Raw: v}
}
