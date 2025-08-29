package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/netbox"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
)

func main() {
	n := netbox.NewEnvironment()
	f := constant.Format

	for _, t := range n.Tenants() {
		fmt.Printf("Tenant: %s\n", t.Format(f))
	}

	for _, s := range n.Sites() {
		fmt.Printf("Site: %s\n", s.Format(f))
	}

	for _, r := range n.DeviceRoles() {
		fmt.Printf("DeviceRole: %s\n", r.Format(f))
	}

	for _, m := range n.Manufacturers() {
		fmt.Printf("Manufacturer: %s\n", m.Format(f))
	}

	for _, t := range n.DeviceTypes() {
		fmt.Printf("DeviceType: %s\n", t.Format(f))
	}

	for _, d := range n.Devices() {
		fmt.Printf("Device: %s\n", d.Format(f))
	}
}
