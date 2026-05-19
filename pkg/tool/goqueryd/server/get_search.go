package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetSearch(
	w http.ResponseWriter,
	_ *http.Request,
	v server.GetSearchParams,
) {
	limit := 10

	if v.Limit != nil {
		limit = *v.Limit
	}

	collection := ""

	if v.Collection != nil {
		collection = *v.Collection
	}

	mode := "hybrid"

	if v.Mode != nil {
		mode = string(*v.Mode)
	}

	sourceType := ""

	if v.SourceType != nil {
		sourceType = *v.SourceType
	}

	outcome := s.service.Search(
		v.Query,
		limit,
		collection,
		v.Full != nil && *v.Full,
		sourceType,
		mode,
	)
	web.EncodeNotation(w, outcome.Results)
}
