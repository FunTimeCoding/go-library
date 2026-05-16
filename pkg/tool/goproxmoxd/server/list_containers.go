package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListContainers(
	w http.ResponseWriter,
	_ *http.Request,
	v server.ListContainersParams,
) {
	var result []*server.Container

	if v.Node != nil && *v.Node != "" {
		result = convert.Containers(
			s.client.MustContainers(s.client.MustNode(*v.Node)),
		)
	} else {
		for _, ns := range s.client.MustNodes() {
			result = append(
				result,
				convert.Containers(
					s.client.MustContainers(s.client.MustNode(ns.Node)),
				)...,
			)
		}
	}

	web.EncodeNotation(w, result)
}
