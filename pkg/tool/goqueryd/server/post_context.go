package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) PostContext(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.PostContextJSONRequestBody
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	s.service.AddContext(body.Collection, body.PathPrefix, body.Description)
	web.EncodeNotation(w, map[string]string{"status": "ok"})
}
