package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) DeleteContainerSnapshot(
	w http.ResponseWriter,
	_ *http.Request,
	identifier int64,
	name string,
	v server.DeleteContainerSnapshotParams,
) {
	c := s.findContainer(identifier, v.Node)

	if c == nil {
		w.WriteHeader(http.StatusNotFound)

		return
	}

	web.EncodeNotation(
		w,
		server.TaskResult{
			TaskId: string(
				s.client.MustDeleteContainerSnapshot(c, name).UPID,
			),
		},
	)
}
