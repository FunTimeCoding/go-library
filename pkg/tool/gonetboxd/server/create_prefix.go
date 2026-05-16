package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/netbox/site"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreatePrefix(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreatePrefixRequest

	if e := json.NewDecoder(q.Body).Decode(&body); e != nil {
		web.InvalidRequestBody(w)

		return
	}

	var si *site.Site

	if body.Site != nil && *body.Site != "" {
		found, e := s.client.SiteByName(*body.Site)

		if e != nil {
			s.captureDetail(w, e)

			return
		}

		si = found
	}

	var description string

	if body.Description != nil {
		description = *body.Description
	}

	p, e := s.client.CreatePrefix(body.Prefix, si, description)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(w, convert.Prefix(p))
}
