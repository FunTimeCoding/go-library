package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) GetTags(
	_ context.Context,
	_ server.GetTagsRequestObject,
) (server.GetTagsResponseObject, error) {
	tags := convert.Tags(s.habitica.MustTags())
	result := make(server.GetTags200JSONResponse, len(tags))

	for i, t := range tags {
		result[i] = *t
	}

	return result, nil
}
