package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateDeviceType(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body generated.CreateDeviceTypeRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	m := s.client.MustManufacturerByName(body.Manufacturer)
	t := s.client.MustCreateDeviceType(body.Model, m)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		generated.DeviceType{
			Identifier:   t.Identifier,
			Model:        t.Model,
			Manufacturer: &m.Name,
		},
	)
}
