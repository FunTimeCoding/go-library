package git

import (
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"strings"
)

func ParseProject(path string) (string, string) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	count := len(parts)

	if count != 2 {
		unexpected.Integer(count)
	}

	namespace := parts[0]
	repository := strings.TrimSuffix(parts[count-1], Directory)

	return namespace, repository
}
