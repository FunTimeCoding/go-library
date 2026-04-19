package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"strings"
)

func playerDetailTable(rows []store.ProfessionRow) g.Node {
	if len(rows) == 0 {
		return h.P(h.Em(g.Text("No stats for this player.")))
	}

	return h.Table(
		h.THead(
			h.Tr(
				h.Th(g.Text("Profession")),
				h.Th(g.Text("Fights")),
				h.Th(g.Text("Damage")),
				h.Th(g.Text("DPS")),
				h.Th(g.Text("Deaths")),
				h.Th(g.Text("Strips")),
			),
		),
		h.TBody(
			g.Map(
				rows,
				func(r store.ProfessionRow) g.Node {
					dps := 0

					if r.ActiveTimeMS > 0 {
						dps = r.Damage * 1000 / r.ActiveTimeMS
					}

					icon := fmt.Sprintf(
						"/static/icons/%s.png",
						strings.ToLower(r.Profession),
					)

					return h.Tr(
						h.Td(
							h.Img(
								h.Src(icon),
								h.Width("20"),
								h.Height("20"),
							),
							g.Textf(" %s", r.Profession),
						),
						h.Td(g.Textf("%d", r.Fights)),
						h.Td(g.Textf("%d", r.Damage)),
						h.Td(g.Textf("%d", dps)),
						h.Td(g.Textf("%d", r.DeadCount)),
						h.Td(g.Textf("%d", r.BoonStrips)),
					)
				},
			),
		),
	)
}
