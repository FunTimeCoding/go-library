package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListNetworks(
	w http.ResponseWriter,
	_ *http.Request,
	name string,
) {
	node := s.client.MustNode(name)
	networks := s.client.MustNetworks(node)
	web.EncodeNotation(w, convert.Networks(networks))
}
