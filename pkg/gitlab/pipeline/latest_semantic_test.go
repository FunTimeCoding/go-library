package pipeline

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"gitlab.com/gitlab-org/api/client-go"
	"testing"
)

func TestLatestSemantic(t *testing.T) {
	assert.Any(
		t,
		&gitlab.PipelineInfo{Ref: "v1.0.2"},
		LatestSemantic(
			[]*gitlab.PipelineInfo{
				{Ref: "v1.0.0"},
				{Ref: "v1.0.2"},
				{Ref: "v1.0.1"},
			},
		),
	)
}
