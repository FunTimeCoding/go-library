package issue

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/google/go-github/v59/github"
	"testing"
)

func TestIssue(t *testing.T) {
	i := New(
		&github.Issue{
			RepositoryURL: ptr.To(
				"https://api.github.com/repos/funtimecoding/go-library",
			),
			Title:   ptr.To(strings.Alfa),
			HTMLURL: ptr.To(strings.Bravo),
		},
	)
	i.Raw = nil
	assert.Any(
		t,
		&Issue{
			Repository: "funtimecoding/go-library",
			Title:      strings.Alfa,
			Link:       strings.Bravo,
		},
		i,
	)
}
