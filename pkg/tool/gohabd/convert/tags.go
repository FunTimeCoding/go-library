package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica"
	"github.com/funtimecoding/go-library/pkg/tool/gohabd/server"
)

func Tags(v []habitica.Tag) []server.Tag {
	result := make([]server.Tag, 0, len(v))

	for _, t := range v {
		result = append(result, Tag(t))
	}

	return result
}
