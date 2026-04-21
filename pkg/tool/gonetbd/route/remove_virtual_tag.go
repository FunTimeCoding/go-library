package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) RemoveVirtualTag(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
	tag string,
) {
	vm := h.client.RemoveVirtualTag(name, tag)
	entry := server.VirtualMachine{Identifier: vm.Identifier, Name: vm.Name}

	if len(vm.Tags) > 0 {
		entry.Tags = &vm.Tags
	}

	web.EncodeNotation(w, entry)
}
