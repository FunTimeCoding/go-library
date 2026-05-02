package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func raidsTable(rows []store.RaidRow) gomponents.Node {
	if len(rows) == 0 {
		return html.P(html.Em(gomponents.Text("No raids created yet.")))
	}

	return html.Table(
		html.THead(
			html.Tr(
				html.Th(gomponents.Text("Name")),
				html.Th(gomponents.Text("Date")),
				html.Th(gomponents.Text("Fights")),
				html.Th(gomponents.Text("Players")),
			),
		),
		html.TBody(
			gomponents.Map(
				rows,
				func(r store.RaidRow) gomponents.Node {
					return html.Tr(
						html.Td(
							html.A(
								html.Href(
									fmt.Sprintf("/raids/%d", r.ID),
								),
								gomponents.Text(r.Name),
							),
						),
						html.Td(
							gomponents.Text(r.Date.Format("2006-01-02")),
						),
						html.Td(gomponents.Textf("%d", r.Fights)),
						html.Td(gomponents.Textf("%d", r.Players)),
					)
				},
			),
		),
	)
}
