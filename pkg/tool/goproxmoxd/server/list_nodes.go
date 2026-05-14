package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListNodes(
	w http.ResponseWriter,
	_ *http.Request,
) {
	nodes := s.client.MustNodes()
	web.EncodeNotation(w, convert.Nodes(nodes))
}
