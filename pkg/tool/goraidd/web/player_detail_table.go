package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func playerDetailTable(rows []store.PlayerRaidRow) g.Node {
	if len(rows) == 0 {
		return h.P(h.Em(g.Text("No stats for this player.")))
	}

	grouped := groupByRaid(rows)
	var nodes []g.Node

	for _, group := range grouped {
		first := group[0]
		nodes = append(
			nodes,
			h.H3(
				h.A(
					h.Href(
						fmt.Sprintf("/raids/%d", first.RaidID),
					),
					g.Text(first.RaidName),
				),
				g.Textf(" — %d fights", first.RaidFights),
			),
			playerRaidTable(group),
		)
	}

	return g.Group(nodes)
}
