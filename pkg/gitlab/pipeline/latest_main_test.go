package pipeline

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/git/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestLatestMain(t *testing.T) {
	assert.Any(
		t,
		&gitlab.PipelineInfo{Ref: "main", SHA: "Bravo"},
		LatestMain(
			[]*gitlab.PipelineInfo{
				{Ref: constant.MainBranch, SHA: strings.Alfa},
				{Ref: constant.MainBranch, SHA: strings.Bravo},
				{Ref: constant.MainBranch, SHA: strings.Charlie},
			},
			strings.Bravo,
		),
	)
}
