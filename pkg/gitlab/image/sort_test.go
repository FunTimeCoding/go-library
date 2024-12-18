package image

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	gitlab "gitlab.com/gitlab-org/api/client-go"
	"testing"
)

func TestSort(t *testing.T) {
	assert.Any(
		t,
		[]*gitlab.RegistryRepositoryTag{
			{Path: "stub:1.0.2"},
			{Path: "stub:1.0.1"},
			{Path: "stub:1.0.0"},
		},
		Sort(
			[]*gitlab.RegistryRepositoryTag{
				{Path: "stub:1.0.0"},
				{Path: "stub:1.0.1"},
				{Path: "stub:1.0.2"},
			},
		),
	)
}
