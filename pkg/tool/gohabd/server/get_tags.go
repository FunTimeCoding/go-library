package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetTags(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(w, convert.Tags(s.habitica.Tags()))
}
