package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) PostCollection(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.PostCollectionJSONRequestBody
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	pattern := ""

	if body.Pattern != nil {
		pattern = *body.Pattern
	}

	s.store.AddCollection(body.Name, body.Path, pattern)
	web.EncodeNotation(w, map[string]string{"status": "ok"})
}
