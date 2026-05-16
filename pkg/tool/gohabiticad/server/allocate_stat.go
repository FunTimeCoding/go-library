package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) AllocateStat(
	w http.ResponseWriter,
	_ *http.Request,
	t server.AllocateStatParamsStat,
) {
	result, e := s.habitica.Allocate(string(t))

	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)

		return
	}

	web.EncodeNotation(w, convert.Stats(result))
}
