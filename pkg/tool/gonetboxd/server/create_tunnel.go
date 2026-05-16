package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateTunnel(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreateTunnelRequest

	if e := json.NewDecoder(q.Body).Decode(&body); e != nil {
		web.InvalidRequestBody(w)

		return
	}

	g, e := s.client.TunnelGroupByName(body.Group)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	t, f := s.client.CreateTunnel(body.Name, body.Encapsulation, g)

	if f != nil {
		s.captureDetail(w, f)

		return
	}

	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(w, convert.Tunnel(t))
}
