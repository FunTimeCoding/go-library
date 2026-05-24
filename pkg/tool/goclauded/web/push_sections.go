package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/web/layout"
	"net/http"
)

func (s *Server) pushSections(
	w http.ResponseWriter,
	f http.Flusher,
) {
	if items := s.usageSummary(); len(items) > 0 {
		pushEvent(
			w,
			layout.SummaryStrip,
			layout.SummaryStripContent(items),
		)
	}

	pushEvent(w, constant.Roster, s.rosterSection())
	pushEvent(w, constant.Activity, s.activitySection())
	f.Flush()
}
