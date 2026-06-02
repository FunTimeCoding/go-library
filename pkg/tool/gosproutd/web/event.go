package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/go-library/pkg/web/layout"
	"net/http"
)

func (s *Server) event() http.HandlerFunc {
	return layout.HandleServerSideEvent(
		s.service.Notifier(),
		func(
			w http.ResponseWriter,
			f http.Flusher,
		) {
			layout.PushEvent(w, constant.Seeds, s.seedTable())
			f.Flush()
		},
	)
}
