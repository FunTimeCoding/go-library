package branch

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings"
	"gitlab.com/gitlab-org/api/client-go"
	"testing"
)

func TestBranch(t *testing.T) {
	assert.True(
		t,
		New(
			&gitlab.Branch{
				Name:   strings.Alfa,
				Merged: false,
				Commit: &gitlab.Commit{
					CreatedAt: ptr.To(constant.StartOfTime),
				},
			},
		) != nil,
	)
}
