package branch

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestBranch(t *testing.T) {
	assert.NotNil(
		t,
		New(
			&gitlab.Branch{
				Name:   upper.Alfa,
				Merged: false,
				Commit: &gitlab.Commit{
					CreatedAt: new(constant.StartOfTime),
				},
			},
		),
	)
}
