package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) StopMachine(
	w http.ResponseWriter,
	_ *http.Request,
	vmid int64,
	v server.StopMachineParams,
) {
	vm := s.findMachine(vmid, v.Node)

	if vm == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	task := s.client.MustStopMachine(vm)
	web.EncodeNotation(w, server.TaskResult{TaskId: string(task.UPID)})
}
