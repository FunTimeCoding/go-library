package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateCluster(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body generated.CreateClusterRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	t := s.client.ClusterTypeByName(body.Type)
	i := s.client.SiteByName(body.Site)
	cl := s.client.CreateCluster(body.Name, t, i)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		generated.Cluster{
			Identifier: cl.Identifier,
			Name:       cl.Name,
			Type:       &t.Name,
			Site:       &i.Name,
		},
	)
}
