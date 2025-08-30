package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/netbox"
)

func readVirtual(
	n *netbox.Client,
	f *option.Format,
) {
	// Virtualization

	for _, g := range n.ClusterGroups() {
		fmt.Printf("ClusterGroup: %s\n", g.Format(f))
	}

	for _, t := range n.ClusterTypes() {
		fmt.Printf("ClusterType: %s\n", t.Format(f))
	}

	for _, c := range n.Clusters() {
		fmt.Printf("Cluster: %s\n", c.Format(f))
	}

	for _, m := range n.VirtualMachines() {
		fmt.Printf("VirtualMachine: %s\n", m.Format(f))
	}
}
