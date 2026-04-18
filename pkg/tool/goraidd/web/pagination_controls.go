package web

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func paginationControls(offset, total int, startValue, endValue string) g.Node {
	var nodes []g.Node

	if offset > 0 {
		previous := offset - pageSize

		if previous < 0 {
			previous = 0
		}

		nodes = append(
			nodes,
			h.A(
				h.Href(paginationLink(previous, startValue, endValue)),
				g.Text("Previous"),
			),
		)
	}

	nodes = append(
		nodes,
		g.Textf(
			" %d–%d of %d ",
			offset+1,
			min(offset+pageSize, total),
			total,
		),
	)

	if offset+pageSize < total {
		nodes = append(
			nodes,
			h.A(
				h.Href(paginationLink(offset+pageSize, startValue, endValue)),
				g.Text("Next"),
			),
		)
	}

	return h.P(h.Class("pagination"), g.Group(nodes))
}
