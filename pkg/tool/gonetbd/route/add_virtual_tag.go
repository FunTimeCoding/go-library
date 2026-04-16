package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) AddVirtualTag(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
	tag string,
) {
	vm := h.client.AddVirtualTag(name, tag)
	w.Header().Set(constant.ContentType, constant.Object)
	entry := server.VirtualMachine{Identifier: vm.Identifier, Name: vm.Name}

	if len(vm.Tags) > 0 {
		entry.Tags = &vm.Tags
	}

	errors.PanicOnError(json.NewEncoder(w).Encode(entry))
}
