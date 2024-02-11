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
	count := len(references)
	var latest *gitlab.PipelineInfo
	index := count - 1

	for _, element := range v {
		if element.Ref == references[index] {
			latest = element
		}
	}

	return latest
}
