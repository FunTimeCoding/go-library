package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func playersTable(rows []store.AttendanceRow) g.Node {
	if len(rows) == 0 {
		return h.P(h.Em(g.Text("No player data yet.")))
	}

	return h.Table(
		h.THead(
			h.Tr(
				h.Th(g.Text("Account")),
				h.Th(g.Text("Characters")),
				h.Th(g.Text("Fights")),
				h.Th(g.Text("Attendance")),
			),
		),
		h.TBody(
			g.Map(
				rows,
				func(r store.AttendanceRow) g.Node {
					pct := 0

					if r.Available > 0 {
						pct = r.Fights * 100 / r.Available
					}

					return h.Tr(
						h.Td(
							h.A(
								h.Href(
									fmt.Sprintf(
										"/players/%s",
										r.Account,
									),
								),
								g.Text(r.Account),
							),
						),
						h.Td(g.Text(r.Characters)),
						h.Td(
							g.Textf("%d / %d", r.Fights, r.Available),
						),
						h.Td(g.Textf("%d%%", pct)),
					)
				},
			),
		),
	)
}
