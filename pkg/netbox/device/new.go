package device

import (
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/netbox/helper"
	"github.com/funtimecoding/go-library/pkg/netbox/tag"
	"github.com/netbox-community/go-netbox/v4"
)

func New(d *netbox.DeviceWithConfigContext) *Device {
	var address string
	var name string
	var comment string

	if d.PrimaryIp4.IsSet() {
		address = d.PrimaryIp4.Get().GetDisplay()
	} else {
		address = constant.NoPrimaryAddress
	}

	if d.Name.IsSet() {
		name = d.GetName()
	} else {
		name = constant.NoName
	}

	if d.Comments == nil {
		comment = constant.NoComment
	} else {
		comment = d.GetComments()
	}

	return &Device{
		Identifier:     d.GetId(),
		Name:           name,
		Comment:        comment,
		PrimaryAddress: address,
		Tags:           tag.Names(d.Tags),
		Link:           helper.ToWebLink(d.GetUrl()),
		Raw:            d,
	}
}
