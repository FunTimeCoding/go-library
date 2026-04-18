package web

import (
	"github.com/funtimecoding/go-library/pkg/gw2"
	gw2Constant "github.com/funtimecoding/go-library/pkg/gw2/constant"
	"github.com/funtimecoding/go-library/pkg/gw2/log_manager/log"
	"github.com/funtimecoding/go-library/pkg/system"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"net/http"
	"slices"
	"time"
)

func (s *Server) logs(
	w http.ResponseWriter,
	r *http.Request,
) {
	all := log.NewSlice(
		gw2.ParseLogs(
			system.ReadBytes(s.logCachePath, gw2Constant.LogFile),
			false,
		),
	)
	startValue := r.URL.Query().Get("start")
	endValue := r.URL.Query().Get("end")

	if startValue == "" {
		startValue = time.Now().Format("2006-01-02") + "T00:00"
	}

	var filtered []*log.Log

	for _, l := range all {
		if startValue != "" {
			start, e := time.Parse("2006-01-02T15:04", startValue)

			if e == nil && l.Time.Before(start) {
				continue
			}
		}

		if endValue != "" {
			end, e := time.Parse("2006-01-02T15:04", endValue)

			if e == nil && l.Time.After(end) {
				continue
			}
		}

		filtered = append(filtered, l)
	}

	slices.Reverse(filtered)
	total := len(filtered)
	offset := parseIntParameter(r, "offset", 0)

	if offset > total {
		offset = total
	}

	end := offset + pageSize

	if end > total {
		end = total
	}

	page := filtered[offset:end]

	if s.isHTMX(r) {
		renderFragment(w, logsTable(page, offset, total, startValue, endValue))

		return
	}

	renderPage(
		w,
		layout(
			"Logs",
			"/",
			h.H1(g.Textf("Combat Logs (%d)", total)),
			dateFilter(startValue, endValue),
			h.Form(
				h.Class("generate-form"),
				h.Method("post"),
				h.Action("/generate"),
				logsTable(page, offset, total, startValue, endValue),
				h.Button(
					h.Type("submit"),
					g.Text("Generate Report"),
				),
			),
		),
	)
}
