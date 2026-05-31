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
	identifier int64,
	v server.GetContainerParams,
) {
	if v.Node != nil && *v.Node != "" {
		web.EncodeNotation(
			w,
			convert.ContainerDetail(
				s.client.MustContainer(
					s.client.MustNode(*v.Node),
					int(identifier),
				),
			),
		)

		return
	}

	for _, n := range s.client.MustNodes() {
		o := s.client.MustNode(n.Node)

		for _, listed := range s.client.MustContainers(o) {
			if int64(listed.VMID) == identifier {
				c := s.client.MustContainer(o, int(identifier))
				web.EncodeNotation(w, convert.ContainerDetail(c))

				return
			}
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
