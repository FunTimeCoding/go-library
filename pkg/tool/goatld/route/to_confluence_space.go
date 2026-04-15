package route

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/server"
)

func toConfluenceSpaces(spaces []*space.Space) []server.ConfluenceSpace {
	result := make([]server.ConfluenceSpace, 0, len(spaces))

	for _, s := range spaces {
		entry := server.ConfluenceSpace{
			Identifier: s.Identifier,
			Key:        s.Raw.Key,
			Name:       s.Name,
		}

		if s.Raw.Type != "" {
			entry.Type = &s.Raw.Type
		}

		if s.Raw.Description != "" {
			entry.Description = &s.Raw.Description
		}

		if s.Raw.HomepageId != "" {
			entry.HomepageIdentifier = &s.Raw.HomepageId
		}

		result = append(result, entry)
	}

	return result
}
