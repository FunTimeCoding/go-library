package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/network"
)

func Physical() {
	n := internal.NetBox()

	for _, p := range n.PhysicalAddressesByHardware(
		network.PhysicalAddress(constant.PhysicalTest0),
	) {
		fmt.Printf("Read physical address: %+v\n", p)
	}
}
