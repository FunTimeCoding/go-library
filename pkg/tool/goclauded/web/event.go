package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/web/layout"
	"net/http"
)

func (s *Server) event() http.HandlerFunc {
	return layout.HandleServerSideEvent(
		s.notifier,
		func(
			w http.ResponseWriter,
			f http.Flusher,
		) {
			if items := s.usageSummary(); len(items) > 0 {
				layout.PushEvent(
					w,
					layout.SummaryStrip,
					layout.SummaryStripContent(items),
				)
			}

			layout.PushEvent(w, constant.Roster, s.rosterSection())
			layout.PushEvent(w, constant.Activity, s.activitySection(nil))
			f.Flush()
		},
	)
}
