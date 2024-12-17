package pipeline

import (
	"gitlab.com/gitlab-org/api/client-go"
	"golang.org/x/mod/semver"
)

func Latest(v []*gitlab.PipelineInfo) *gitlab.PipelineInfo {
	result := v[0]

	for _, element := range v {
		if result == nil || semver.Compare(element.Ref, result.Ref) > 0 {
			result = element
		}
	}

	return result
}
