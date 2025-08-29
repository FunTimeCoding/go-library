package tag

import "github.com/netbox-community/go-netbox/v4"

type Tag struct {
	Identifier int32
	Name       string

	Raw    *netbox.Tag
	Nested *netbox.NestedTagRequest
}
