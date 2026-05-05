package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetGear(
	w http.ResponseWriter,
	_ *http.Request,
) {
	result, e := s.habitica.UserGear()

	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)

		return
	}

	web.EncodeNotation(w, convert.Gear(result))
}
