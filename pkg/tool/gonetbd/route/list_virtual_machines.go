package route

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (h *Router) ListVirtualMachines(
	w http.ResponseWriter,
	_ *http.Request,
) {
	machines := h.client.VirtualMachines()
	result := make([]server.VirtualMachine, 0, len(machines))

	for _, m := range machines {
		entry := server.VirtualMachine{Identifier: m.Identifier, Name: m.Name}

		if m.Raw.Cluster.IsSet() && m.Raw.Cluster.Get() != nil {
			n := m.Raw.Cluster.Get().GetName()
			entry.Cluster = &n
		}

		if m.Raw.Site.IsSet() && m.Raw.Site.Get() != nil {
			n := m.Raw.Site.Get().GetName()
			entry.Site = &n
		}

		if len(m.Tags) > 0 {
			entry.Tags = &m.Tags
		}

		result = append(result, entry)
	}

	web.EncodeNotation(w, result)
}
