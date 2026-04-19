package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/store"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
	"strings"
)

func raidDetailTable(rows []store.RaidPlayerRow) g.Node {
	if len(rows) == 0 {
		return h.P(h.Em(g.Text("No player stats for this raid.")))
	}

	return h.Table(
		h.THead(
			h.Tr(
				h.Th(g.Text("Player")),
				h.Th(g.Text("Profession")),
				h.Th(g.Text("Fights")),
				h.Th(g.Text("Dmg/s")),
				h.Th(g.Text("Heal/s")),
				h.Th(g.Text("Clean/s")),
				h.Th(g.Text("Strip/s")),
				h.Th(g.Text("Barrier/s")),
				h.Th(g.Text("Downs/s")),
				h.Th(g.Text("Death/m")),
				h.Th(g.Text("Tag Dist")),
			),
		),
		h.TBody(
			g.Map(
				rows,
				func(r store.RaidPlayerRow) g.Node {
					seconds := float64(r.ActiveTimeMS) / 1000
					minutes := seconds / 60
					icon := fmt.Sprintf(
						"/static/icons/%s.png",
						strings.ToLower(r.Profession),
					)

					return h.Tr(
						h.Td(
							h.A(
								h.Href(
									fmt.Sprintf(
										"/players/%s",
										r.Account,
									),
								),
								g.Text(r.Name),
							),
						),
						h.Td(
							h.Img(
								h.Src(icon),
								h.Width("20"),
								h.Height("20"),
							),
							g.Textf(" %s", r.Profession),
						),
						h.Td(g.Textf("%d", r.Fights)),
						h.Td(g.Text(perSecond(r.Damage, seconds, 1))),
						h.Td(g.Text(perSecond(r.Healing, seconds, 1))),
						h.Td(
							g.Text(
								perSecond(r.ConditionCleanses, seconds, 1),
							),
						),
						h.Td(g.Text(perSecond(r.BoonStrips, seconds, 2))),
						h.Td(g.Text(perSecond(r.Barrier, seconds, 1))),
						h.Td(g.Text(perSecond(r.Downs, seconds, 1))),
						h.Td(g.Text(perMinute(r.DeadCount, minutes))),
						h.Td(g.Textf("%.0f", r.DistToCom)),
					)
				},
			),
		),
	)
}
