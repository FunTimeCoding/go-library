package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
)

func (s *Server) GetSearch(
	_ context.Context,
	r server.GetSearchRequestObject,
) (server.GetSearchResponseObject, error) {
	limit := 10

	if r.Params.Limit != nil {
		limit = *r.Params.Limit
	}

	collection := ""

	if r.Params.Collection != nil {
		collection = *r.Params.Collection
	}

	mode := "hybrid"

	if r.Params.Mode != nil {
		mode = string(*r.Params.Mode)
	}

	var metadata map[string]string

	if r.Params.Metadata != nil {
		metadata = *r.Params.Metadata
	}

	if r.Params.SourceType != nil && *r.Params.SourceType != "" {
		if metadata == nil {
			metadata = map[string]string{}
		}

		metadata[constant.SourceType] = *r.Params.SourceType
	}

	outcome := s.service.Search(
		r.Params.Query,
		limit,
		collection,
		r.Params.Full != nil && *r.Params.Full,
		mode,
		metadata,
	)
	result := server.GetSearch200JSONResponse{
		Results: convertSearchResults(outcome.Results),
	}

	if outcome.Degraded {
		result.Degraded = new(bool(true))
	}

	if len(outcome.Facets) > 0 {
		facets := convertFacets(outcome.Facets)
		result.Facets = &facets
	}

	return result, nil
}
