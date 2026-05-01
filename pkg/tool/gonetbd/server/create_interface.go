package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	upstream "github.com/netbox-community/go-netbox/v4"
	"net/http"
)

func (s *Server) CreateInterface(
	w http.ResponseWriter,
	q *http.Request,
	name string,
) {
	var body generated.CreateInterfaceRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	d := s.client.MustDeviceByNameStrict(name)
	i := s.client.MustCreateInterface(
		d,
		body.Name,
		upstream.InterfaceTypeValue(body.Type),
	)
	result := generated.Interface{
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
