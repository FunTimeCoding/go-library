package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) EquipGear(
	w http.ResponseWriter,
	_ *http.Request,
	key string,
) {
	e := s.habitica.Equip(key)

	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)

		return
	}

	result, f := s.habitica.UserGear()

	if f != nil {
		http.Error(w, f.Error(), http.StatusInternalServerError)

		return
	}

	web.EncodeNotation(w, convert.Gear(result))
}
