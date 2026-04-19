package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func raidsTable(rows []store.RaidRow) g.Node {
	if len(rows) == 0 {
		return h.P(h.Em(g.Text("No raids created yet.")))
	}

	return h.Table(
		h.THead(
			h.Tr(
				h.Th(g.Text("Name")),
				h.Th(g.Text("Date")),
				h.Th(g.Text("Fights")),
				h.Th(g.Text("Players")),
			),
		),
		h.TBody(
			g.Map(
				rows,
				func(r store.RaidRow) g.Node {
					return h.Tr(
						h.Td(
							h.A(
								h.Href(
									fmt.Sprintf("/raids/%d", r.ID),
								),
								g.Text(r.Name),
							),
						),
						h.Td(
							g.Text(r.Date.Format("2006-01-02")),
						),
						h.Td(g.Textf("%d", r.Fights)),
						h.Td(g.Textf("%d", r.Players)),
					)
				},
			),
		),
	)
}
