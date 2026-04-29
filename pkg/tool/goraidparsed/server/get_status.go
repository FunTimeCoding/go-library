package server

import (
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidparsed/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetStatus(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(
		w,
		generated.StatusResponse{Status: "ok"},
	)
}
