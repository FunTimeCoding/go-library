package server

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetList(
	w http.ResponseWriter,
	_ *http.Request,
	v server.GetListParams,
) {
	documents, e := s.service.ListDocuments(v.Collection)
	errors.PanicOnError(e)
	web.EncodeNotation(w, documents)
}
