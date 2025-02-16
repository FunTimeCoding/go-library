package issue

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestIssue(t *testing.T) {
	actual := New(
		&sentry.Issue{
			Project:   &sentry.Project{Name: strings.Alfa},
			Type:      ptr.To(strings.Bravo),
			Title:     ptr.To(strings.Charlie),
			Permalink: ptr.To(strings.Delta),
		},
	)
	actual.Raw = nil
	assert.Any(
		t,
		&Issue{
			Project: "Alfa",
			Type:    "Bravo",
			Title:   "Charlie",
			Link:    "Delta",
		},
		actual,
	)
}
