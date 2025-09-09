package pipeline

import (
	"gitlab.com/gitlab-org/api/client-go"
	"log"
)

func LatestMain(
	v []*gitlab.PipelineInfo,
	mainHash string,
) *gitlab.PipelineInfo {
	for _, e := range v {
		if e.SHA == mainHash {
			return e
		}
	}

	log.Panicf("main hash not found for %s", mainHash)

	return &gitlab.PipelineInfo{}
}
