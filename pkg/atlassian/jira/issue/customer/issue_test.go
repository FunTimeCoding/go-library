package customer

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestIssue(t *testing.T) {
	actual := New(
		&models.CustomerRequestScheme{
			IssueKey: strings.Alfa,
			RequestFieldValues: []*models.CustomerRequestRequestFieldValueScheme{
				{FieldID: SummaryField, Value: strings.Bravo},
				{FieldID: DescriptionField, Value: strings.Charlie},
			},
			CurrentStatus: &models.CustomerRequestCurrentStatusScheme{
				Status: constant.ServiceDeskResolved,
			},
			Links: &models.CustomerRequestLinksScheme{
				Web: strings.Delta,
			},
		},
	)
	actual.Raw = nil
	assert.Any(
		t,
		&Issue{
			Key:    "Alfa",
			Status: "Resolved",
			Title:  "Bravo",
			Body:   "Charlie",
			Link:   "Delta",
		},
		actual,
	)
}
