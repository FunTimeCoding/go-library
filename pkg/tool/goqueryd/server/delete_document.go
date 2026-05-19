package server

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) DeleteDocument(
	w http.ResponseWriter,
	_ *http.Request,
	v server.DeleteDocumentParams,
) {
	deleted, e := s.store.DeleteDocument(v.Collection, v.Path)
	errors.PanicOnError(e)
	web.EncodeNotation(w, map[string]bool{"deleted": deleted})
}
