package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreatePage(
	w http.ResponseWriter,
	e *http.Request,
) {
	var c server.CreatePageJSONRequestBody
	errors.PanicOnError(json.NewDecoder(e.Body).Decode(&s))
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		convert.ConfluencePageDetail(
			s.confluence.MustCreatePage(
				c.SpaceIdentifier,
				c.ParentIdentifier,
				c.Title,
				c.Body,
			),
		),
	)
}
