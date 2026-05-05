package option

import "github.com/funtimecoding/go-library/pkg/netbox"

func New() *Netbox {
	return &Netbox{Client: netbox.NewEnvironment()}
}
