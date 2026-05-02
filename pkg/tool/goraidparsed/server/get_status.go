package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goraidparsed/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetStatus(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(
		w,
		server.StatusResponse{Status: "ok"},
	)
}
