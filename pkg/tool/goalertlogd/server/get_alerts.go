package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetAlerts(
	w http.ResponseWriter,
	_ *http.Request,
	v server.GetAlertsParams,
) {
	web.EncodeNotation(w, toResponse(s.store.ByName(v.Name)))
}
