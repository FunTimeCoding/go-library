package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gw2/log_manager/log"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func logsTable(logs []*log.Log) g.Node {
	if len(logs) == 0 {
		return h.P(h.Em(g.Text("No logs found.")))
	}

	return h.Table(
		h.THead(
			h.Tr(
				h.Th(),
				h.Th(g.Text("Time")),
				h.Th(g.Text("Duration")),
				h.Th(g.Text("Map")),
				h.Th(g.Text("Players")),
			),
		),
		h.TBody(
			g.Map(
				logs,
				func(l *log.Log) g.Node {
					return h.Tr(
						h.Td(
							h.Input(
								h.Type("checkbox"),
								h.Name("fileNames"),
								h.Value(l.Raw.FileName),
							),
						),
						h.Td(g.Text(
							l.Time.Format("2006-01-02 15:04"),
						)),
						h.Td(g.Text(l.Raw.EncounterDuration)),
						h.Td(g.Text(
							fmt.Sprintf("%d", l.Raw.MapID),
						)),
						h.Td(g.Text(
							fmt.Sprintf("%d", len(l.Raw.Players)),
						)),
					)
				},
			),
		),
	)
}
