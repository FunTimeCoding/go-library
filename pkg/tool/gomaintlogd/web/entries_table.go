package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	g "maragu.dev/gomponents"
	hx "maragu.dev/gomponents-htmx"
	h "maragu.dev/gomponents/html"
)

func entriesTable(entries []store.Entry) g.Node {
	if len(entries) == 0 {
		return h.P(h.Em(g.Text("No entries found.")))
	}

	var rows []g.Node

	for _, e := range entries {
		target := fmt.Sprintf("detail-%d", e.ID)
		rows = append(
			rows,
			h.Tr(
				h.ID(fmt.Sprintf("row-%d", e.ID)),
				h.Class("clickable-row"),
				hx.Get(fmt.Sprintf("/entry/detail?id=%d", e.ID)),
				hx.Target(fmt.Sprintf("#%s", target)),
				hx.Swap("outerHTML"),
				h.Td(g.Text(formatTime(e.Timestamp))),
				h.Td(g.Text(e.Action)),
				h.Td(g.Text(e.User)),
				h.Td(g.Text(e.System)),
				h.Td(g.Text(e.Service)),
				h.Td(g.Text(truncate(e.Description, 80))),
			),
			h.Tr(
				h.ID(target),
				h.Style("display:none"),
			),
		)
	}

	return h.Table(
		h.THead(
			h.Tr(
				h.Th(g.Text("Timestamp")),
				h.Th(g.Text("Action")),
				h.Th(g.Text("User")),
				h.Th(g.Text("System")),
				h.Th(g.Text("Service")),
				h.Th(g.Text("Description")),
			),
		),
		h.TBody(g.Group(rows)),
	)
}
