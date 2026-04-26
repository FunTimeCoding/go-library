package issue

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestIssue(t *testing.T) {
	r := response.NewIssue()
	r.ID = "1"
	r.Type = strings.Bravo
	r.Title = strings.Charlie
	r.Permalink = strings.Delta
	r.Project.Name = strings.Alfa
	actual := New(r)
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
