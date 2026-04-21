package option

import "github.com/funtimecoding/go-library/pkg/netbox"

type Netbox struct {
	Client *netbox.Client
	Port   int
}
