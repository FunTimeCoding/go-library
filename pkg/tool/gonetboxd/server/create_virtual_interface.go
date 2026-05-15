package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateVirtualInterface(
	w http.ResponseWriter,
	q *http.Request,
	name string,
) {
	var body server.CreateVirtualInterfaceRequest

	if e := json.NewDecoder(q.Body).Decode(&body); e != nil {
		web.InvalidRequestBody(w)

		return
	}

	vm, e := s.client.VirtualMachineByName(name)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	i, f := s.client.CreateVirtualInterface(vm, body.Name)

	if f != nil {
		s.captureDetail(w, f)

		return
	}

	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		server.VirtualInterface{
			Identifier: i.GetId(), Name: i.GetName(),
		},
	)
}
