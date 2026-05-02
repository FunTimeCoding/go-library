package web

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"os"
	"strings"
)

func (s *Server) reports(
	w http.ResponseWriter,
	r *http.Request,
) {
	var reports []os.DirEntry

	for _, n := range system.ReadDirectory(s.outputPath) {
		if !n.IsDir() && strings.HasSuffix(
			n.Name(),
			constant.HypertextExtension,
		) {
			reports = append(reports, n)
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
			html.H1(gomponents.Text("Generated Reports")),
			reportsTable(reports),
		),
	)
}
