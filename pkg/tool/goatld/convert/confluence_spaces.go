package convert

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/generated/server"
)

func ConfluenceSpaces(v []*space.Space) []*server.ConfluenceSpace {
	result := make([]*server.ConfluenceSpace, 0, len(v))

	for _, s := range v {
		p := &server.ConfluenceSpace{
			Identifier: s.Identifier,
			Key:        s.Raw.Key,
			Name:       s.Name,
		}

		if s.Raw.Type != "" {
			p.Type = &s.Raw.Type
		}

		if s.Raw.Description != "" {
			p.Description = &s.Raw.Description
		}

		if s.Raw.HomepageId != "" {
			p.HomepageIdentifier = &s.Raw.HomepageId
		}

		result = append(result, p)
	}

	return result
}
