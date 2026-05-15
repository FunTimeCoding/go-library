package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateDevice(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreateDeviceRequest

	if e := json.NewDecoder(q.Body).Decode(&body); e != nil {
		web.InvalidRequestBody(w)

		return
	}

	role, e := s.client.DeviceRoleByName(body.Role)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	deviceType, f := s.client.DeviceTypeByName(body.Type)

	if f != nil {
		s.captureDetail(w, f)

		return
	}

	site, g := s.client.SiteByName(body.Site)

	if g != nil {
		s.captureDetail(w, g)

		return
	}

	var tags []string

	if body.Tags != nil {
		tags = *body.Tags
	}

	var ten *tenant.Tenant

	if body.Tenant != nil && *body.Tenant != "" {
		t, h := s.client.TenantByName(*body.Tenant)

		if h != nil {
			s.captureDetail(w, h)

			return
		}

		ten = t
	}

	d, i := s.client.CreateDevice(
		body.Name,
		role,
		tags,
		deviceType,
		site,
		ten,
	)

	if i != nil {
		s.captureDetail(w, i)

		return
	}

	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(w, convert.Device(d))
}
