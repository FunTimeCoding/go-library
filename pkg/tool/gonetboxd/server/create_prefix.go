package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
)

func (s *Server) CreatePrefix(
	_ context.Context,
	r server.CreatePrefixRequestObject,
) (server.CreatePrefixResponseObject, error) {
	var si *site.Site

	if r.Body.Site != nil && *r.Body.Site != "" {
		found, e := s.client.SiteByName(*r.Body.Site)

		if e != nil {
			return server.CreatePrefix500JSONResponse(*s.captureDetail(e)), nil
		}

		si = found
	}

	var description string

	if r.Body.Description != nil {
		description = *r.Body.Description
	}

	p, e := s.client.CreatePrefix(r.Body.Prefix, si, description)

	if e != nil {
		return server.CreatePrefix500JSONResponse(*s.captureDetail(e)), nil
	}

	return server.CreatePrefix201JSONResponse(*convert.Prefix(p)), nil
}
