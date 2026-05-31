package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) DeleteMachineSnapshot(
	w http.ResponseWriter,
	_ *http.Request,
	identifier int64,
	name string,
	v server.DeleteMachineSnapshotParams,
) {
	m := s.findMachine(identifier, v.Node)

	if m == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	web.EncodeNotation(
		w,
		server.TaskResult{
			TaskId: string(
				s.client.MustDeleteMachineSnapshot(m, name).UPID,
			),
		},
	)
}
