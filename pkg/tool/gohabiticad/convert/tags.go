package convert

import (
	"github.com/funtimecoding/go-library/pkg/habitica/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func Tags(v []response.Tag) []server.Tag {
	result := make([]server.Tag, 0, len(v))

	for _, t := range v {
		result = append(result, Tag(t))
	}

	return result
}
