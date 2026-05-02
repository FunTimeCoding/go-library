package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateDeviceRole(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreateNameRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	r := s.client.MustCreateDeviceRole(body.Name)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		server.DeviceRole{
			Identifier: r.Identifier, Name: r.Name,
		},
	)
}
