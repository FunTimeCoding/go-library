package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetContainer(
	w http.ResponseWriter,
	_ *http.Request,
	vmid int64,
	v server.GetContainerParams,
) {
	if v.Node != nil && *v.Node != "" {
		node := s.client.MustNode(*v.Node)
		c := s.client.MustContainer(node, int(vmid))
		web.EncodeNotation(w, convert.ContainerDetail(c))

		return
	}

	for _, ns := range s.client.MustNodes() {
		node := s.client.MustNode(ns.Node)

		for _, listed := range s.client.MustContainers(node) {
			if int64(listed.VMID) == vmid {
				c := s.client.MustContainer(node, int(vmid))
				web.EncodeNotation(w, convert.ContainerDetail(c))

				return
			}
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
