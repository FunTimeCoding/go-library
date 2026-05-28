package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) PostTag(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.PostTagJSONRequestBody
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	collection := ""

	if body.Collection != nil {
		collection = *body.Collection
	}

	s.service.SetSourceType(collection, body.Path, body.SourceType)
	web.EncodeNotation(w, map[string]string{"status": "ok"})
}
