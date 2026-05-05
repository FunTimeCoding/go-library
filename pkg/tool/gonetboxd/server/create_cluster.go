package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateCluster(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreateClusterRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	t := s.client.MustClusterTypeByName(body.Type)
	i := s.client.MustSiteByName(body.Site)
	cl := s.client.MustCreateCluster(body.Name, t, i)
	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(
		w,
		server.Cluster{
			Identifier: cl.Identifier,
			Name:       cl.Name,
			Type:       &t.Name,
			Site:       &i.Name,
		},
	)
}
