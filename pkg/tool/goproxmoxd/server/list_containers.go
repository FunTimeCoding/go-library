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
	var result []server.Container

	if v.Node != nil && *v.Node != "" {
		node := s.client.MustNode(*v.Node)
		containers := s.client.MustContainers(node)
		result = convert.Containers(containers)
	} else {
		for _, ns := range s.client.MustNodes() {
			node := s.client.MustNode(ns.Node)
			containers := s.client.MustContainers(node)
			result = append(result, convert.Containers(containers)...)
		}
	}

	web.EncodeNotation(w, result)
}
