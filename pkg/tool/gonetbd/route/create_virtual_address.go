package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) CreateVirtualAddress(
	w http.ResponseWriter,
	q *http.Request,
	name string,
) {
	var body generated.CreateAddressRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	vm := h.client.VirtualMachineByName(name)
	i := h.client.VirtualMachineInterfaceByName(vm, body.Interface)
	a := h.client.CreateVirtualAddress(i.GetId(), body.Address)
	w.Header().Set(constant.ContentType, constant.Object)
	w.WriteHeader(http.StatusCreated)
	errors.PanicOnError(
		json.NewEncoder(w).Encode(
			generated.Address{
				Identifier: a.Identifier, Address: a.Address.String(),
			},
		),
	)
}
