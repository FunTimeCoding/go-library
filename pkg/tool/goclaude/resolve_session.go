package goclaude

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"strings"
)

func resolveSession(
	c *client.ClientWithResponses,
	query string,
) string {
	response, e := c.GetSessionsWithResponse(
		context.Background(),
		&client.GetSessionsParams{},
	)

	if e != nil || response.JSON200 == nil {
		return ""
	}

	for _, s := range response.JSON200.Sessions {
		if s.Alias != nil && *s.Alias == query {
			return s.Identifier
		}
	}

	for _, s := range response.JSON200.Sessions {
		if s.Identifier == query {
			return s.Identifier
		}

		if s.Slug != nil && *s.Slug == query {
			return s.Identifier
		}

		if strings.HasPrefix(s.Identifier, query) {
			return s.Identifier
		}
	}

	return ""
}
