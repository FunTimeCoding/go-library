package network

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/system/constant"
)

func (i *Interface) Format(f *option.Format) string {
	s := status.New(f).String(i.formatName(f), i.formatType()).RawList(i.Raw)

	if i.PhysicalAddress.String() != constant.NullPhysicalAddress.String() {
		s.String(i.PhysicalAddress.String())
	}

	s.String(i.Description)

	for _, element := range i.Contexts {
		s.Line("  Context: %+v", element)
	}

	for _, element := range i.VirtualNetworks {
		s.Line("  Virtual network: %+v", element)
	}

	for _, element := range i.WirelessNetworks {
		s.Line("  Wireless network: %+v", element)
	}

	return s.Format()
}
