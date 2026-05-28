package server

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetDocument(
	w http.ResponseWriter,
	_ *http.Request,
	v server.GetDocumentParams,
) {
	document, _, e := s.service.GetDocument(v.Path)
	errors.PanicOnError(e)
	web.EncodeNotation(w, document)
}
