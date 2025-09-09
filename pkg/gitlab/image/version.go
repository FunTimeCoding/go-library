package image

import (
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
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
		return key_value.Empty(constant.VersionPrefix, result)
	}

	return result
}
