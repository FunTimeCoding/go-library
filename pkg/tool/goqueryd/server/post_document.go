package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) PostDocument(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.PostDocumentJSONRequestBody
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	metadata := map[string]string{}

	if body.Metadata != nil {
		for key, value := range *body.Metadata {
			metadata[key] = value
		}
	}

	if body.SourceType != nil && *body.SourceType != "" {
		metadata[constant.SourceType] = *body.SourceType
	}

	errors.PanicOnError(
		s.service.PushDocument(
			body.Collection,
			body.Path,
			body.Body,
			metadata,
		),
	)
	web.EncodeNotation(w, map[string]string{"status": "ok"})
}
