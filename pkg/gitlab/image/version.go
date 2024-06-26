package image

import (
	"fmt"
	"github.com/xanzy/go-gitlab"
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
