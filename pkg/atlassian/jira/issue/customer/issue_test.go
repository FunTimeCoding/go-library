package customer

import (
	"github.com/ctreminiom/go-atlassian/v2/pkg/infra/models"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"testing"
)

func TestIssue(t *testing.T) {
	actual := New(
		&models.CustomerRequestScheme{
			IssueKey: upper.Alfa,
			RequestFieldValues: []*models.CustomerRequestRequestFieldValueScheme{
				{FieldID: SummaryField, Value: upper.Bravo},
				{FieldID: DescriptionField, Value: upper.Charlie},
			},
			CurrentStatus: &models.CustomerRequestCurrentStatusScheme{
				Status: constant.ServiceDeskResolved,
			},
			Links: &models.CustomerRequestLinksScheme{
				Web: upper.Delta,
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
