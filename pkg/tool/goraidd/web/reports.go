package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"net/http"
	"os"
	"strings"
)

func (s *Server) reports(
	w http.ResponseWriter,
	r *http.Request,
) {
	entries, e := os.ReadDir(s.outputPath)
	errors.PanicOnError(e)
	var reports []os.DirEntry

	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".html") {
			reports = append(reports, entry)
		}
	}

	if s.isHTMX(r) {
		renderFragment(w, reportsTable(reports))

		return
	}

	renderPage(
		w,
		layout(
			"Reports",
			"/reports",
			h.H1(g.Text("Generated Reports")),
			reportsTable(reports),
		),
	)
}
