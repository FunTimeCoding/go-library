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
	vmid int64,
	v server.GetMachineParams,
) {
	if v.Node != nil && *v.Node != "" {
		node := s.client.MustNode(*v.Node)
		vm := s.client.MustMachine(node, int(vmid))
		web.EncodeNotation(w, convert.MachineDetail(vm))

		return
	}

	for _, ns := range s.client.MustNodes() {
		node := s.client.MustNode(ns.Node)

		for _, listed := range s.client.MustMachines(node) {
			if int64(listed.VMID) == vmid {
				vm := s.client.MustMachine(node, int(vmid))
				web.EncodeNotation(w, convert.MachineDetail(vm))

				return
			}
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
