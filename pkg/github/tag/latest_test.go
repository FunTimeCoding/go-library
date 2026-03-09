package tag

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/google/go-github/v84/github"
	"testing"
)

func TestLatest(t *testing.T) {
	assert.String(
		t,
		"v1.0.1",
		Latest(
			[]*github.RepositoryTag{
				{Name: new("v1.0.0")},
				{Name: new("v1.0.1")},
			},
		).GetName(),
	)
}
