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
	var result []*server.Machine

	if v.Node != nil && *v.Node != "" {
		result = convert.Machines(
			s.client.MustMachines(s.client.MustNode(*v.Node)),
		)
	} else {
		for _, ns := range s.client.MustNodes() {
			result = append(
				result,
				convert.Machines(
					s.client.MustMachines(s.client.MustNode(ns.Node)),
				)...,
			)
		}
	}

	web.EncodeNotation(w, result)
}
