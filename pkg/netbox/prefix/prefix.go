package prefix

import "github.com/netbox-community/go-netbox/v4"

type Prefix struct {
	Identifier  int32
	Name        string
	Description string

	Raw *netbox.Prefix
}
