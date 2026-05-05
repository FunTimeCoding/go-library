package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
)

func Tenant(t *tenant.Tenant) server.Tenant {
	return server.Tenant{Identifier: t.Identifier, Name: t.Name}
}
