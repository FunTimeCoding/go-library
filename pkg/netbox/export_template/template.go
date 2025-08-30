package export_template

import "github.com/netbox-community/go-netbox/v4"

type Template struct {
	Identifier int32
	Name       string

	Raw *netbox.ExportTemplate
}
