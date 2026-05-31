package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetMachine(
	w http.ResponseWriter,
	_ *http.Request,
	identifier int64,
	v server.GetMachineParams,
) {
	if v.Node != nil && *v.Node != "" {
		web.EncodeNotation(
			w,
			convert.MachineDetail(
				s.client.MustMachine(
					s.client.MustNode(*v.Node),
					int(identifier),
				),
			),
		)

		return
	}

	for _, n := range s.client.MustNodes() {
		o := s.client.MustNode(n.Node)

		for _, listed := range s.client.MustMachines(o) {
			if int64(listed.VMID) == identifier {
				web.EncodeNotation(
					w,
					convert.MachineDetail(
						s.client.MustMachine(o, int(identifier)),
					),
				)

				return
			}
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
