package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/netbox-community/go-netbox/v4"
	"net/http"
)

func (s *Server) CreateInterface(
	w http.ResponseWriter,
	q *http.Request,
	name string,
) {
	var body server.CreateInterfaceRequest

	if e := json.NewDecoder(q.Body).Decode(&body); e != nil {
		web.InvalidRequestBody(w)

		return
	}

	d, e := s.client.DeviceByNameStrict(name)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	i, f := s.client.CreateInterface(
		d,
		body.Name,
		netbox.InterfaceTypeValue(body.Type),
	)

	if f != nil {
		s.captureDetail(w, f)

		return
	}

	result := server.Interface{
		Identifier: i.Identifier,
		Name:       i.Name,
	}

	if i.Type != "" {
		result.Type = new(string(i.Type))
	}

	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(w, result)
}
