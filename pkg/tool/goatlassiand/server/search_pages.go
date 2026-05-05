package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) SearchPages(
	w http.ResponseWriter,
	_ *http.Request,
	p server.SearchPagesParams,
) {
	web.EncodeNotation(
		w,
		convert.ConfluencePages(s.confluence.MustSearch(p.Query)),
	)
}
