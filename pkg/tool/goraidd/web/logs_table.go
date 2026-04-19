package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/raid"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func logsTable(
	fights []raid.Fight,
	offset, total int,
	startValue, endValue string,
	filtered bool,
) g.Node {
	if len(fights) == 0 {
		return h.P(h.Em(g.Text("No logs found.")))
	}

	var nodes []g.Node

	if !filtered {
		nodes = append(
			nodes,
			paginationControls(offset, total, startValue, endValue),
		)
	}

	nodes = append(
		nodes,
		h.Table(
			h.THead(
				h.Tr(
					h.Th(),
					h.Th(g.Text("Time")),
					h.Th(g.Text("Duration")),
					h.Th(g.Text("Map")),
					h.Th(g.Text("Players")),
					h.Th(g.Text("Enemies")),
				),
			),
			h.TBody(
				g.Map(
					fights,
					func(f raid.Fight) g.Node {
						return h.Tr(
							h.Td(
								h.Input(
									h.Type("checkbox"),
									h.Name("fileNames"),
									h.Value(f.Filename),
								),
							),
							h.Td(
								g.Text(
									f.Timestamp.Format("2006-01-02 15:04"),
								),
							),
							h.Td(g.Text(fightDuration(f))),
							h.Td(g.Text(fightMap(f))),
							h.Td(
								g.Text(
									fmt.Sprintf("%d", f.AlliedCount),
								),
							),
							h.Td(g.Text(fightEnemies(f))),
						)
					},
				),
			),
		),
	)

	if !filtered {
		nodes = append(
			nodes,
			paginationControls(offset, total, startValue, endValue),
		)
	}

	return g.Group(nodes)
}
