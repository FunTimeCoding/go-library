package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/netbox"
)

func readTunnel(
	n *netbox.Client,
	f *option.Format,
) {
	// VPN

	for _, g := range n.TunnelGroups() {
		fmt.Printf("TunnelGroup: %s\n", g.Format(f))
	}

	for _, t := range n.Tunnels() {
		fmt.Printf("Tunnel: %s\n", t.Format(f))
	}
}
