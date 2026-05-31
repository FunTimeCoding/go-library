package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateMachineSnapshot(
	w http.ResponseWriter,
	r *http.Request,
	identifier int64,
	v server.CreateMachineSnapshotParams,
) {
	m := s.findMachine(identifier, v.Node)

	if m == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	var b server.SnapshotRequest
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&b))
	web.EncodeNotation(
		w,
		server.TaskResult{
			TaskId: string(
				s.client.MustCreateMachineSnapshot(m, b.Name).UPID,
			),
		},
	)
}
