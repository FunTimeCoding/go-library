package pipeline

import "gitlab.com/gitlab-org/api/client-go/v2"

func LatestMain(
	v []*gitlab.PipelineInfo,
	mainHash string,
) *gitlab.PipelineInfo {
	for _, e := range v {
		if e.SHA == mainHash {
			return e
		}
	}

	return nil
}
