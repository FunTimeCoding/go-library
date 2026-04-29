package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateDevice(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body generated.CreateDeviceRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	role := s.client.DeviceRoleByName(body.Role)
	deviceType := s.client.DeviceTypeByName(body.Type)
	site := s.client.SiteByName(body.Site)
	var tags []string

	if body.Tags != nil {
		tags = *body.Tags
	}

	var ten *tenant.Tenant

	if body.Tenant != nil && *body.Tenant != "" {
		ten = s.client.TenantByName(*body.Tenant)
	}

	d := s.client.CreateDevice(body.Name, role, tags, deviceType, site, ten)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(w, toDevice(d))
}
