package network

import (
	"github.com/netbox-community/go-netbox/v4"
	"log"
	"slices"
)

func validateType(
	interfaceTypes []netbox.InterfaceTypeValue,
	t netbox.InterfaceTypeValue,
) {
	if !slices.Contains(interfaceTypes, t) {
		log.Panicf("unexpected interface type: %s", t)
	}
}
