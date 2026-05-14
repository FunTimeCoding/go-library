package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListMachines(
	w http.ResponseWriter,
	_ *http.Request,
	v server.ListMachinesParams,
) {
	var result []server.Machine

	if v.Node != nil && *v.Node != "" {
		node := s.client.MustNode(*v.Node)
		machines := s.client.MustMachines(node)
		result = convert.Machines(machines)
	} else {
		for _, ns := range s.client.MustNodes() {
			node := s.client.MustNode(ns.Node)
			machines := s.client.MustMachines(node)
			result = append(result, convert.Machines(machines)...)
		}
	}

	web.EncodeNotation(w, result)
}
