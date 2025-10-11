package issue

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"github.com/google/go-github/v70/github"
	"testing"
)

func TestIssue(t *testing.T) {
	i := New(
		&github.Issue{
			RepositoryURL: locator.New(
				"api.github.com",
			).Path("/repos/funtimecoding/go-library").Pointer(),
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
