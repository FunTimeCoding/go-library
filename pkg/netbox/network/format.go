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

	for _, c := range i.Contexts {
		s.Line("  Context: %+v", c)
	}

	for _, n := range i.VirtualNetworks {
		s.Line("  Virtual network: %+v", n)
	}

	for _, n := range i.WirelessNetworks {
		s.Line("  Wireless network: %+v", n)
	}

	return s.Format()
}
