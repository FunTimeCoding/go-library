package issue

import (
	"github.com/atlassian/go-sentry-api"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
	"time"
)

func TestIssue(t *testing.T) {
	actual := New(
		&sentry.Issue{
			ID:        ptr.To("1"),
			Project:   &sentry.Project{Name: strings.Alfa},
			Type:      ptr.To(strings.Bravo),
			Title:     ptr.To(strings.Charlie),
			Permalink: ptr.To(strings.Delta),
			FirstSeen: ptr.To(
				time.Date(
					2020,
					1,
					1,
					0,
					0,
					0,
					0,
					time.UTC,
				),
			),
		},
	)
	actual.Create = time.Time{}
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
