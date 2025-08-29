package netbox

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/netbox/cache"
	"github.com/netbox-community/go-netbox/v4"
)

type Client struct {
	context context.Context
	client  *netbox.APIClient
	cache   *cache.Cache

	interfaceTypes []netbox.InterfaceTypeValue

	verbose bool
}
