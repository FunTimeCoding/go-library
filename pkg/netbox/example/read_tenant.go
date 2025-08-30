package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/netbox"
)

func readTenant(
	n *netbox.Client,
	f *option.Format,
) {
	for _, g := range n.TenantGroups() {
		fmt.Printf("TenantGroup: %s\n", g.Format(f))
	}

	for _, t := range n.Tenants() {
		fmt.Printf("Tenant: %s\n", t.Format(f))
	}

	for _, g := range n.ContactGroups() {
		fmt.Printf("ContactGroup: %s\n", g.Format(f))
	}

	for _, r := range n.ContactRoles() {
		fmt.Printf("ContactRole: %s\n", r.Format(f))
	}

	for _, c := range n.Contacts() {
		fmt.Printf("Contact: %s\n", c.Format(f))
	}
}
