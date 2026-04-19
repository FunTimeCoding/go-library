package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"os"
)

func reportsTable(reports []os.DirEntry) g.Node {
	if len(reports) == 0 {
		return h.P(h.Em(g.Text("No reports generated yet.")))
	}

	return h.Table(
		h.THead(
			h.Tr(
				h.Th(g.Text("Report")),
				h.Th(g.Text("Size")),
				h.Th(),
			),
		),
		h.TBody(
			g.Map(
				reports,
				func(entry os.DirEntry) g.Node {
					info, e := entry.Info()
					errors.PanicOnError(e)

					return h.Tr(
						h.Td(
							h.A(
								h.Href(
									fmt.Sprintf(
										"/reports/%s",
										entry.Name(),
									),
								),
								g.Text(entry.Name()),
							),
						),
						h.Td(
							h.Class("report-size"),
							g.Text(formatSize(info.Size())),
						),
						h.Td(
							h.Form(
								h.Method("post"),
								h.Action(
									fmt.Sprintf(
										"/reports/%s/delete",
										entry.Name(),
									),
								),
								h.Button(
									h.Type("submit"),
									h.Class("secondary outline"),
									g.Text("Delete"),
								),
							),
						),
					)
				},
			),
		),
	)
}
