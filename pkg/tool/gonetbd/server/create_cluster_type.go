package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateClusterType(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body generated.CreateNameRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	t := s.client.CreateClusterType(body.Name)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		generated.ClusterType{
			Identifier: t.Identifier,
			Name:       t.Name,
		},
	)
}
