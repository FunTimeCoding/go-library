package web

import (
	"github.com/funtimecoding/go-library/pkg/raid"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"net/http"
	"time"
)

func (s *Server) logs(
	w http.ResponseWriter,
	r *http.Request,
) {
	all := s.store.Fights()
	filtered := r.URL.Query().Has("start") || r.URL.Query().Has("end")
	startValue := r.URL.Query().Get("start")
	endValue := r.URL.Query().Get("end")

	if !filtered {
		startValue = time.Now().Format("2006-01-02") + "T00:00"
		endValue = time.Now().Format("2006-01-02") + "T23:59"
	}

	var fights []raid.Fight

	if filtered {
		for _, f := range all {
			if startValue != "" {
				start, e := time.Parse("2006-01-02T15:04", startValue)

				if e == nil && f.Timestamp.Before(start) {
					continue
				}
			}

			if endValue != "" {
				end, e := time.Parse("2006-01-02T15:04", endValue)

				if e == nil && f.Timestamp.After(end) {
					continue
				}
			}

			fights = append(fights, f)
		}
	} else {
		if len(all) > pageSize {
			fights = all[:pageSize]
		} else {
			fights = all
		}
	}

	total := len(fights)

	if filtered {
		if s.isHTMX(r) {
			renderFragment(
				w,
				logsTable(fights, 0, total, startValue, endValue, true),
			)

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
					logsTable(fights, 0, total, startValue, endValue, true),
					h.Button(
						h.Type("submit"),
						h.FormAction("/generate"),
						g.Text("Generate Report"),
					),
					h.Button(
						h.Type("submit"),
						h.FormAction("/raids/create"),
						h.Class("secondary"),
						g.Text("Create Raid"),
					),
				),
			),
		)

		return
	}

	offset := parseIntParameter(r, "offset", 0)

	if offset > total {
		offset = total
	}

	end := offset + pageSize

	if end > total {
		end = total
	}

	page := fights[offset:end]

	if s.isHTMX(r) {
		renderFragment(
			w,
			logsTable(page, offset, total, startValue, endValue, false),
		)

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
				logsTable(page, offset, total, startValue, endValue, false),
				h.Button(
					h.Type("submit"),
					h.FormAction("/generate"),
					g.Text("Generate Report"),
				),
				h.Button(
					h.Type("submit"),
					h.FormAction("/raids/create"),
					h.Class("secondary"),
					g.Text("Create Raid"),
				),
			),
		),
	)
}
