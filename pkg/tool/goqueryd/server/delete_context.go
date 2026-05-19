package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) DeleteContext(
	w http.ResponseWriter,
	_ *http.Request,
	v server.DeleteContextParams,
) {
	web.EncodeNotation(
		w,
		map[string]bool{
			"deleted": s.store.RemoveContext(v.Collection, v.PathPrefix),
		},
	)
}
