package network

import (
	"github.com/netbox-community/go-netbox/v4"
	"log"
	"slices"
)

func validateType(t netbox.InterfaceTypeValue) {
	if !slices.Contains(Interfaces, t) {
		log.Panicf("unexpected interface type: %s", t)
	}
}
