package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetPage(
	w http.ResponseWriter,
	_ *http.Request,
	identifier string,
) {
	web.EncodeNotation(
		w,
		convert.ConfluencePageDetail(s.confluence.MustPage(identifier)),
	)
}
