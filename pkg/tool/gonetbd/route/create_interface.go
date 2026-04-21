package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	upstream "github.com/netbox-community/go-netbox/v4"
	"net/http"
)

func (h *Router) CreateInterface(
	w http.ResponseWriter,
	q *http.Request,
	name string,
) {
	var body generated.CreateInterfaceRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	d := h.client.DeviceByNameStrict(name)
	i := h.client.CreateInterface(
		d,
		body.Name,
		upstream.InterfaceTypeValue(body.Type),
	)
	result := generated.Interface{
		Identifier: i.Identifier,
		Name:       i.Name,
	}

	if i.Type != "" {
		s := string(i.Type)
		result.Type = &s
	}

	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(w, result)
}
