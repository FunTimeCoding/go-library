package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateAddress(
	w http.ResponseWriter,
	q *http.Request,
	name string,
) {
	var body generated.CreateAddressRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	d := s.client.MustDeviceByNameStrict(name)
	i := s.client.MustDeviceInterfaceByNameStrict(d, body.Interface)
	a := s.client.MustCreateAddress(i.Identifier, body.Address)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		generated.Address{
			Identifier: a.Identifier,
			Address:    a.Address.String(),
		},
	)
}
