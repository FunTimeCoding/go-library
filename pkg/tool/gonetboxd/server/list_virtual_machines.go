package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListVirtualMachines(
	w http.ResponseWriter,
	_ *http.Request,
) {
	machines, e := s.client.VirtualMachines()

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	result := make([]server.VirtualMachine, 0, len(machines))

	for _, m := range machines {
		entry := server.VirtualMachine{Identifier: m.Identifier, Name: m.Name}

		if m.Raw.Cluster.IsSet() && m.Raw.Cluster.Get() != nil {
			entry.Cluster = new(m.Raw.Cluster.Get().GetName())
		}

		if m.Raw.Site.IsSet() && m.Raw.Site.Get() != nil {
			entry.Site = new(m.Raw.Site.Get().GetName())
		}

		if len(m.Tags) > 0 {
			entry.Tags = &m.Tags
		}

		result = append(result, entry)
	}

	web.EncodeNotation(w, result)
}
