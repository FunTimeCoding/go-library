package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) PostImpressions(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.PostImpressionsJSONBody

	if json.NewDecoder(r.Body).Decode(&body) != nil {
		web.InvalidRequestBody(w)

		return
	}

	source := ""

	if body.Source != nil {
		source = *body.Source
	}

	identifier, e := s.service.CreateImpression(body.Content, source)

	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)

		return
	}

	web.EncodeNotation(
		w,
		server.ImpressionResponse{Identifier: int(identifier)},
	)
}
