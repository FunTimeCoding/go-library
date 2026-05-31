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
	web.EncodeNotation(
		w,
		convert.Networks(s.client.MustNetworks(s.client.MustNode(name))),
	)
}
