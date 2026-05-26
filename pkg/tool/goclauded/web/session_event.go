package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/web/layout"
	"net/http"
)

func (s *Server) sessionEvent() http.HandlerFunc {
	return layout.HandleServerSideEventWithRequest(
		s.notifier,
		func(
			w http.ResponseWriter,
			f http.Flusher,
			r *http.Request,
		) {
			identifier := r.PathValue(constant.Identifier)
			layout.PushEvent(w, constant.Pulse, s.pulseSection(identifier))
			f.Flush()
		},
	)
}
