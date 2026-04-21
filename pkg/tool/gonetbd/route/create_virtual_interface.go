package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) CreateVirtualInterface(
	w http.ResponseWriter,
	q *http.Request,
	name string,
) {
	var body generated.CreateVirtualInterfaceRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	vm := h.client.VirtualMachineByName(name)
	i := h.client.CreateVirtualInterface(vm, body.Name)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		generated.VirtualInterface{
			Identifier: i.GetId(), Name: i.GetName(),
		},
	)
}
