package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateVirtualMachine(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreateVirtualMachineRequest

	if e := json.NewDecoder(q.Body).Decode(&body); e != nil {
		web.InvalidRequestBody(w)

		return
	}

	cl, e := s.client.ClusterByName(body.Cluster)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	vm, f := s.client.CreateVirtualMachine(body.Name, cl)

	if f != nil {
		s.captureDetail(w, f)

		return
	}

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
