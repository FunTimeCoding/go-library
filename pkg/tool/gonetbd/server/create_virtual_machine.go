package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateVirtualMachine(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body generated.CreateVirtualMachineRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	cl := s.client.ClusterByName(body.Cluster)
	vm := s.client.CreateVirtualMachine(body.Name, cl)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		generated.VirtualMachine{
			Identifier: vm.Identifier,
			Name:       vm.Name,
			Cluster:    &cl.Name,
		},
	)
}
