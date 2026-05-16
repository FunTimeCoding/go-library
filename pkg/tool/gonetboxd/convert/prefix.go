package convert

import (
	"github.com/funtimecoding/go-library/pkg/netbox/prefix"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
)

func Prefix(p *prefix.Prefix) server.Prefix {
	result := server.Prefix{
		Identifier: p.Identifier,
		Prefix:     p.Name,
	}

	if p.Description != "" {
		result.Description = &p.Description
	}

	if p.Raw.ScopeType.IsSet() && p.Raw.ScopeId.IsSet() {
		result.Site = new(p.Raw.GetDisplay())
	}

	return result
}
