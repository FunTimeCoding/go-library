package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/gonetbd/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (h *Router) CreateCluster(w http.ResponseWriter, q *http.Request) {
	var body generated.CreateClusterRequest
	errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
	t := h.client.ClusterTypeByName(body.Type)
	s := h.client.SiteByName(body.Site)
	cl := h.client.CreateCluster(body.Name, t, s)
	w.Header().Set(constant.ContentType, constant.Object)
	w.WriteHeader(http.StatusCreated)
	errors.PanicOnError(json.NewEncoder(w).Encode(generated.Cluster{Identifier: cl.Identifier, Name: cl.Name, Type: &t.Name, Site: &s.Name}))
}
