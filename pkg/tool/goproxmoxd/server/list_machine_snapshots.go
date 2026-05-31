package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListMachineSnapshots(
	w http.ResponseWriter,
	_ *http.Request,
	identifier int64,
	v server.ListMachineSnapshotsParams,
) {
	m := s.findMachine(identifier, v.Node)

	if m == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	web.EncodeNotation(
		w,
		convert.MachineSnapshots(s.client.MustMachineSnapshots(m)),
	)
}
