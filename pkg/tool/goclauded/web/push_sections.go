package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"net/http"
)

func (s *Server) pushSections(
	w http.ResponseWriter,
	flusher http.Flusher,
) {
	pushEvent(w, constant.Roster, s.rosterSection())
	pushEvent(w, constant.Activity, s.activitySection())
	flusher.Flush()
}
