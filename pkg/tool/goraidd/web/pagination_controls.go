package web

import (
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
)

func paginationControls(
	offset, total int,
	startValue, endValue string,
) gomponents.Node {
	var nodes []gomponents.Node

	if offset > 0 {
		previous := offset - pageSize

		if previous < 0 {
			previous = 0
		}

		nodes = append(
			nodes,
			html.A(
				html.Href(paginationLink(previous, startValue, endValue)),
				gomponents.Text("Previous"),
			),
		)
	}

	nodes = append(
		nodes,
		gomponents.Textf(
			" %d–%d of %d ",
			offset+1,
			min(offset+pageSize, total),
			total,
		),
	)

	if offset+pageSize < total {
		nodes = append(
			nodes,
			html.A(
				html.Href(paginationLink(offset+pageSize, startValue, endValue)),
				gomponents.Text("Next"),
			),
		)
	}

	return html.P(html.Class("pagination"), gomponents.Group(nodes))
}
