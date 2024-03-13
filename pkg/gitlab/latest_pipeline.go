package gitlab

import (
	"github.com/xanzy/go-gitlab"
	"golang.org/x/mod/semver"
)

func LatestPipeline(v []*gitlab.PipelineInfo) *gitlab.PipelineInfo {
	var references []string

	for _, element := range v {
		references = append(references, element.Ref)
	}

	semver.Sort(references)
	var result *gitlab.PipelineInfo
	index := len(references) - 1

	for _, element := range v {
		if element.Ref == references[index] {
			result = element
		}
	}

	return result
}
