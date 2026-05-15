package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateDeviceType(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreateDeviceTypeRequest

	if e := json.NewDecoder(q.Body).Decode(&body); e != nil {
		web.InvalidRequestBody(w)

		return
	}

	m, e := s.client.ManufacturerByName(body.Manufacturer)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	t, f := s.client.CreateDeviceType(body.Model, m)

	if f != nil {
		s.captureDetail(w, f)

		return
	}

	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		server.DeviceType{
			Identifier:   t.Identifier,
			Model:        t.Model,
			Manufacturer: &m.Name,
		},
	)
}
