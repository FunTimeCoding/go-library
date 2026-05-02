package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
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
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	d := s.client.MustDeviceByNameStrict(name)
	i := s.client.MustCreateInterface(
		d,
		body.Name,
		netbox.InterfaceTypeValue(body.Type),
	)
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
