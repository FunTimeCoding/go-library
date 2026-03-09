package issue

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestIssue(t *testing.T) {
	actual := New(
		&sentry.Issue{
			ID:        new("1"),
			Project:   &sentry.Project{Name: strings.Alfa},
			Type:      new(strings.Bravo),
			Title:     new(strings.Charlie),
			Permalink: new(strings.Delta),
			FirstSeen: new(constant.StartOfTime),
		},
	)
	actual.Create = nil
	actual.Raw = nil
	assert.Any(
		t,
		&Issue{
			MonitorIdentifier: "sentry-1",
			Project:           "Alfa",
			Type:              "Bravo",
			Title:             "Charlie",
			Link:              "Delta",
		},
		actual,
	)
}
