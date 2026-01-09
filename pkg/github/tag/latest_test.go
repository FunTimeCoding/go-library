package tag

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/google/go-github/v81/github"
	"testing"
)

func TestLatest(t *testing.T) {
	assert.String(
		t,
		"v1.0.1",
		Latest(
			[]*github.RepositoryTag{
				{Name: ptr.To("v1.0.0")},
				{Name: ptr.To("v1.0.1")},
			},
		).GetName(),
	)
}
