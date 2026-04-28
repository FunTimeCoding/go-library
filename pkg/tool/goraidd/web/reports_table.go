package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"os"
)

func reportsTable(v []os.DirEntry) g.Node {
	if len(v) == 0 {
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
				v,
				func(d os.DirEntry) g.Node {
					i, e := d.Info()
					errors.PanicOnError(e)

					return h.Tr(
						h.Td(
							h.A(
								h.Href(
									fmt.Sprintf(
										"/reports/%s",
										d.Name(),
									),
								),
								g.Text(d.Name()),
							),
						),
						h.Td(
							h.Class("report-size"),
							g.Text(formatSize(i.Size())),
						),
						h.Td(
							h.Form(
								h.Method("post"),
								h.Action(
									fmt.Sprintf(
										"/reports/%s/delete",
										d.Name(),
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
