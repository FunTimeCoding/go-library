package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateVirtualAddress(
	w http.ResponseWriter,
	q *http.Request,
	name string,
) {
	var body server.CreateAddressRequest

	if e := json.NewDecoder(q.Body).Decode(&body); e != nil {
		web.InvalidRequestBody(w)

		return
	}

	vm, e := s.client.VirtualMachineByName(name)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	i, f := s.client.VirtualMachineInterfaceByName(vm, body.Interface)

	if f != nil {
		s.captureDetail(w, f)

		return
	}

	a, g := s.client.CreateVirtualAddress(i.GetId(), body.Address)

	if g != nil {
		s.captureDetail(w, g)

		return
	}

	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		server.Address{
			Identifier: a.Identifier, Address: a.Address.String(),
		},
	)
}
