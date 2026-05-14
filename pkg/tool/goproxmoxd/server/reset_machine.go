package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ResetMachine(
	w http.ResponseWriter,
	_ *http.Request,
	vmid int64,
	v server.ResetMachineParams,
) {
	vm := s.findMachine(vmid, v.Node)

	if vm == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	task := s.client.MustResetMachine(vm)
	web.EncodeNotation(w, server.TaskResult{TaskId: string(task.UPID)})
}
