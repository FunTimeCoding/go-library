package device_bay_template

import "github.com/netbox-community/go-netbox/v4"

type Template struct {
	Identifier int32
	Model      string

	Raw *netbox.DeviceBayTemplate
}
