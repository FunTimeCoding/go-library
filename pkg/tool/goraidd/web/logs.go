package web

import (
	"github.com/funtimecoding/go-library/pkg/gw2"
	gw2Constant "github.com/funtimecoding/go-library/pkg/gw2/constant"
	"github.com/funtimecoding/go-library/pkg/gw2/log_manager/log"
	"github.com/funtimecoding/go-library/pkg/system"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) logs(
	w http.ResponseWriter,
	r *http.Request,
) {
	logs := log.NewSlice(
		gw2.ParseLogs(
			system.ReadBytes(s.logCachePath, gw2Constant.LogFile),
			false,
		),
	)

	if s.isHTMX(r) {
		renderFragment(w, logsTable(logs))

		return
	}

	renderPage(
		w,
		layout(
			"Logs",
			"/",
			h.H1(g.Text("Combat Logs")),
			h.Form(
				h.Class("generate-form"),
				h.Method("post"),
				h.Action("/api/generate"),
				logsTable(logs),
				h.Button(
					h.Type("submit"),
					g.Text("Generate Report"),
				),
			),
		),
	)
}
