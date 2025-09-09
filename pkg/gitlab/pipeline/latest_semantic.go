package pipeline

import (
	"gitlab.com/gitlab-org/api/client-go"
	"golang.org/x/mod/semver"
	"log"
)

func LatestSemantic(v []*gitlab.PipelineInfo) *gitlab.PipelineInfo {
	if len(v) == 0 {
		log.Panic("empty slice")
	}

	var result *gitlab.PipelineInfo

	for _, e := range v {
		if !semver.IsValid(e.Ref) {
			continue
		}

		if result == nil {
			result = e

			continue
		}

		if semver.Compare(e.Ref, result.Ref) > 0 {
			result = e
		}
	}

	return result
}
