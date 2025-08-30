package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/netbox"
)

func readUser(
	n *netbox.Client,
	f *option.Format,
) {
	for _, g := range n.UserGroups() {
		fmt.Printf("UserGroup: %s\n", g.Format(f))
	}

	for _, g := range n.Users() {
		fmt.Printf("User: %s\n", g.Format(f))
	}
}
