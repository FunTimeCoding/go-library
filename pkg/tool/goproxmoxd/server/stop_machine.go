package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) StopMachine(
	w http.ResponseWriter,
	_ *http.Request,
	identifier int64,
	v server.StopMachineParams,
) {
	m := s.findMachine(identifier, v.Node)

	if m == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	web.EncodeNotation(
		w,
		server.TaskResult{TaskId: string(s.client.MustStopMachine(m).UPID)},
	)
}
