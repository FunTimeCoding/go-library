package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListContainerSnapshots(
	w http.ResponseWriter,
	_ *http.Request,
	identifier int64,
	v server.ListContainerSnapshotsParams,
) {
	c := s.findContainer(identifier, v.Node)

	if c == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	snapshots := s.client.MustContainerSnapshots(c)
	web.EncodeNotation(w, convert.ContainerSnapshots(snapshots))
}
