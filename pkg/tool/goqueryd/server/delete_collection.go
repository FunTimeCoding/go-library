package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) DeleteCollection(
	w http.ResponseWriter,
	_ *http.Request,
	v server.DeleteCollectionParams,
) {
	web.EncodeNotation(
		w,
		map[string]bool{"deleted": s.service.DeleteCollection(v.Name)},
	)
}
