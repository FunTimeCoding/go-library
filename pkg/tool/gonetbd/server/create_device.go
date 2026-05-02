package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateDevice(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreateDeviceRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	role := s.client.MustDeviceRoleByName(body.Role)
	deviceType := s.client.MustDeviceTypeByName(body.Type)
	site := s.client.MustSiteByName(body.Site)
	var tags []string

	if body.Tags != nil {
		tags = *body.Tags
	}

	var ten *tenant.Tenant

	if body.Tenant != nil && *body.Tenant != "" {
		ten = s.client.MustTenantByName(*body.Tenant)
	}

	d := s.client.MustCreateDevice(
		body.Name,
		role,
		tags,
		deviceType,
		site,
		ten,
	)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(w, convert.Device(d))
}
