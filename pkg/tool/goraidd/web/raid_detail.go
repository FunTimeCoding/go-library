package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
	"strconv"
)

func (s *Server) raidDetail(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier, e := strconv.Atoi(r.PathValue("id"))
	errors.PanicOnError(e)
	fights := s.store.RaidFights(identifier)
	players := s.store.RaidPlayerStats(identifier)
	renderPage(
		w,
		layout(
			"Raid",
			"/raids",
			html.H1(gomponents.Textf("Raid - %d fights", len(fights))),
			raidDetailTable(players),
		),
	)
}
