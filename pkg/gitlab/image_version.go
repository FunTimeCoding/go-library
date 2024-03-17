package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"log"
	"strings"
)

func ImageVersion(v *gitlab.RegistryRepositoryTag) string {
	result := strings.Split(v.Path, ":")[1]

	if result == "" {
		log.Panicf("empty version: %+v", result)
	}

	return result
}
