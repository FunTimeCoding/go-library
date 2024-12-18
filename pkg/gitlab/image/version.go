package image

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"gitlab.com/gitlab-org/api/client-go"
	"log"
	"strings"
)

func Version(v *gitlab.RegistryRepositoryTag) string {
	result := strings.Split(v.Path, ":")[1]

	if result == "" {
		log.Panicf("empty version: %+v", result)
	}

	if !strings.HasPrefix(result, constant.VersionPrefix) {
		return fmt.Sprintf("%s%s", constant.VersionPrefix, result)
	}

	return result
}
