package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/netbox"
)

func readWireless(
	n *netbox.Client,
	f *option.Format,
) {
	for _, g := range n.WirelessNetworkGroups() {
		fmt.Printf("WirelessNetworkGroup: %s\n", g.Format(f))
	}

	for _, e := range n.WirelessNetworks() {
		fmt.Printf("WirelessNetwork: %s\n", e.Format(f))
	}

	if false {
		// TODO: What must devices have to show up in the picker?
		for _, l := range n.WirelessLinks() {
			fmt.Printf("WirelessLink: %s\n", l.Format(f))
		}
	}
}
