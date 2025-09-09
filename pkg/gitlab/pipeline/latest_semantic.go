package pipeline

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
	"golang.org/x/mod/semver"
	"log"
)

func LatestSemantic(v []*gitlab.PipelineInfo) *gitlab.PipelineInfo {
	if len(v) == 0 {
		log.Panic("empty slice")
	}

	result := v[0]

	for _, e := range v {
		errors.PanicOnSemver(e.Ref)
		errors.PanicOnSemver(result.Ref)

		if semver.Compare(e.Ref, result.Ref) > 0 {
			result = e
		}
	}

	return result
}
