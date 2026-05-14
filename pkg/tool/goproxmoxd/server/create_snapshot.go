package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateSnapshot(
	w http.ResponseWriter,
	r *http.Request,
	vmid int64,
	v server.CreateSnapshotParams,
) {
	vm := s.findMachine(vmid, v.Node)

	if vm == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	var body server.SnapshotRequest
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	task := s.client.MustCreateSnapshot(vm, body.Name)
	web.EncodeNotation(w, server.TaskResult{TaskId: string(task.UPID)})
}
