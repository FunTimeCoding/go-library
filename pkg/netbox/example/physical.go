package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/netbox"
	"github.com/funtimecoding/go-library/pkg/network"
)

func Physical() {
	n := netbox.NewEnvironment()

	for _, p := range n.PhysicalAddressesByHardware(
		network.PhysicalAddress(constant.PhysicalTest0),
	) {
		fmt.Printf("Read physical address: %+v\n", p)
	}
}
