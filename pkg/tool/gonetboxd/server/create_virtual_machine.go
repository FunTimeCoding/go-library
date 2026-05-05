package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateVirtualMachine(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreateVirtualMachineRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	cl := s.client.MustClusterByName(body.Cluster)
	vm := s.client.MustCreateVirtualMachine(body.Name, cl)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		server.VirtualMachine{
			Identifier: vm.Identifier,
			Name:       vm.Name,
			Cluster:    &cl.Name,
		},
	)
}
