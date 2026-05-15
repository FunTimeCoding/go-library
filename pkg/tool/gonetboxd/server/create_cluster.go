package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateCluster(
	w http.ResponseWriter,
	q *http.Request,
) {
	var body server.CreateClusterRequest

	if e := json.NewDecoder(q.Body).Decode(&body); e != nil {
		web.InvalidRequestBody(w)

		return
	}

	t, e := s.client.ClusterTypeByName(body.Type)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	i, f := s.client.SiteByName(body.Site)

	if f != nil {
		s.captureDetail(w, f)

		return
	}

	cl, g := s.client.CreateCluster(body.Name, t, i)

	if g != nil {
		s.captureDetail(w, g)

		return
	}

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
