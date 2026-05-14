package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetNodeStatus(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
) {
	result := s.client.MustNodeStatus(name)
	web.EncodeNotation(w, convert.NodeStatus(result))
}
