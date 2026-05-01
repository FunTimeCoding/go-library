package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetPageChildren(
	w http.ResponseWriter,
	_ *http.Request,
	identifier string,
) {
	web.EncodeNotation(
		w,
		convert.ConfluencePagesFromPages(
			s.confluence.MustChildPagesByIdentifier(identifier),
		),
	)
}
