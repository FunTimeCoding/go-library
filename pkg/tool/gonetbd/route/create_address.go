package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) CreateAddress(
	w http.ResponseWriter,
	q *http.Request,
	name string,
) {
	var body generated.CreateAddressRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	d := h.client.DeviceByNameStrict(name)
	i := h.client.DeviceInterfaceByNameStrict(d, body.Interface)
	a := h.client.CreateAddress(i.Identifier, body.Address)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		generated.Address{
			Identifier: a.Identifier,
			Address:    a.Address.String(),
		},
	)
}
