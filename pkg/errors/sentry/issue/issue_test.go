package issue

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestIssue(t *testing.T) {
	actual := New(
		&response.Issue{
			ID:        "1",
			Type:      strings.Bravo,
			Title:     strings.Charlie,
			Permalink: strings.Delta,
			Project:   Project{Name: strings.Alfa},
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
