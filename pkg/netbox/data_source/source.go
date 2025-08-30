package data_source

import "github.com/netbox-community/go-netbox/v4"

type Source struct {
	Identifier int32
	Name       string

	Raw *netbox.DataSource
}
