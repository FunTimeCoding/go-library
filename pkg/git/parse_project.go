package git

import (
	"log"
	"strings"
)

func ParseProject(path string) (string, string) {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	count := len(parts)

	if count != 2 {
		log.Panicf("unexpected: %d", count)
	}

	namespace := parts[0]
	repository := strings.TrimSuffix(parts[count-1], Directory)

	return namespace, repository
}
