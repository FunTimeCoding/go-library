package internal

import (
	"github.com/funtimecoding/go-library/pkg/netbox"
	"github.com/funtimecoding/go-library/pkg/netbox/network"
)

func NetBox() *netbox.Client {
	return netbox.NewEnvironment(
		netbox.WithInterfaceTypes(network.Interfaces),
	)
}
