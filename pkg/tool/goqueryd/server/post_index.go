package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) PostIndex(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.PostIndexJSONRequestBody
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	collection := ""

	if body.Collection != nil {
		collection = *body.Collection
	}

	results, e := s.service.IndexCollections(collection)
	errors.PanicOnError(e)
	web.EncodeNotation(w, results)
}
