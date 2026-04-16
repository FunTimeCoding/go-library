package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/netbox/tenant"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) CreateDevice(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body generated.CreateDeviceRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	role := h.client.DeviceRoleByName(body.Role)
	deviceType := h.client.DeviceTypeByName(body.Type)
	site := h.client.SiteByName(body.Site)
	var tags []string

	if body.Tags != nil {
		tags = *body.Tags
	}

	var ten *tenant.Tenant

	if body.Tenant != nil && *body.Tenant != "" {
		ten = h.client.TenantByName(*body.Tenant)
	}

	d := h.client.CreateDevice(body.Name, role, tags, deviceType, site, ten)
	w.Header().Set(constant.ContentType, constant.Object)
	w.WriteHeader(http.StatusCreated)
	errors.PanicOnError(json.NewEncoder(w).Encode(toDevice(d)))
}
