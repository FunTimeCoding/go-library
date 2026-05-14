package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListSnapshots(
	w http.ResponseWriter,
	_ *http.Request,
	vmid int64,
	v server.ListSnapshotsParams,
) {
	vm := s.findMachine(vmid, v.Node)

	if vm == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	snapshots := s.client.MustSnapshots(vm)
	web.EncodeNotation(w, convert.Snapshots(snapshots))
}
