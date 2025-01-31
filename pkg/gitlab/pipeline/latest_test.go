package pipeline

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	gitlab "gitlab.com/gitlab-org/api/client-go"
	"testing"
)

func TestLatest(t *testing.T) {
	assert.Any(
		t,
		&gitlab.PipelineInfo{Ref: "1.0.0"},
		Latest([]*gitlab.PipelineInfo{{Ref: "1.0.0"}}),
	)
}
