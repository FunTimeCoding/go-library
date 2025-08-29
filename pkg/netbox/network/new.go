package network

import (
	"github.com/funtimecoding/go-library/pkg/network"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func New(
	i *netbox.Interface,
	interfaceTypes []netbox.InterfaceTypeValue,
) *Interface {
	var name string

	if s := i.GetName(); s != "" {
		name = i.GetName()
	} else {
		name = NoName
	}

	var h net.HardwareAddr

	if s := i.GetMacAddress(); s != "" {
		h = network.PhysicalAddress(s)
	} else {
		h = constant.NullPhysicalAddress
	}

	var t netbox.InterfaceTypeValue

	if i.Type.Value != nil {
		t = *i.Type.Value
		validateType(interfaceTypes, t)
	}

	return &Interface{
		Identifier:       i.GetId(),
		Name:             name,
		Description:      i.GetDescription(),
		Type:             t,
		PhysicalAddress:  h,
		Contexts:         i.GetVdcs(),
		VirtualNetworks:  i.GetTaggedVlans(),
		WirelessNetworks: i.GetWirelessLans(),
		Raw:              i,
	}
}
