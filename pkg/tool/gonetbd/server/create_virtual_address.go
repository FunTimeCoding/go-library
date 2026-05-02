package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateVirtualAddress(
	w http.ResponseWriter,
	q *http.Request,
	name string,
) {
	var body server.CreateAddressRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	vm := s.client.MustVirtualMachineByName(name)
	i := s.client.MustVirtualMachineInterfaceByName(vm, body.Interface)
	a := s.client.MustCreateVirtualAddress(i.GetId(), body.Address)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		server.Address{
			Identifier: a.Identifier, Address: a.Address.String(),
		},
	)
}
