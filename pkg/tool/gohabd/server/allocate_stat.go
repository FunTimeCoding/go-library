package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) AllocateStat(
	w http.ResponseWriter,
	_ *http.Request,
	stat server.AllocateStatParamsStat,
) {
	result, e := s.habitica.Allocate(string(stat))

	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)

		return
	}

	web.EncodeNotation(w, convert.Stats(result))
}
