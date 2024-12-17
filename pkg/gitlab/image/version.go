package image

import (
	"fmt"
	"gitlab.com/gitlab-org/api/client-go"
	"log"
	"strings"
)

func Version(v *gitlab.RegistryRepositoryTag) string {
	result := strings.Split(v.Path, ":")[1]

	if result == "" {
		log.Panicf("empty version: %+v", result)
	}

	return fmt.Sprintf("v%s", result)
}
