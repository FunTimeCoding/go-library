package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
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

	var metadata map[string]string

	if v.Metadata != nil {
		metadata = *v.Metadata
	}

	if v.SourceType != nil && *v.SourceType != "" {
		if metadata == nil {
			metadata = map[string]string{}
		}

		metadata[constant.SourceType] = *v.SourceType
	}

	outcome := s.service.Search(
		v.Query,
		limit,
		collection,
		v.Full != nil && *v.Full,
		mode,
		metadata,
	)
	web.EncodeNotation(w, outcome)
}
