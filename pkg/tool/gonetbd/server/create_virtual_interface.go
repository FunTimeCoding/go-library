package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateVirtualInterface(
	w http.ResponseWriter,
	q *http.Request,
	name string,
) {
	var body server.CreateVirtualInterfaceRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	vm := s.client.MustVirtualMachineByName(name)
	i := s.client.MustCreateVirtualInterface(vm, body.Name)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		server.VirtualInterface{
			Identifier: i.GetId(), Name: i.GetName(),
		},
	)
}
