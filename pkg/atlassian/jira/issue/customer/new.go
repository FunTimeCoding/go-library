package customer

import (
	"fmt"
	"github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

func New(v *models.CustomerRequestScheme) *Issue {
	var summary string
	var description string

	for _, e := range v.RequestFieldValues {
		switch e.FieldID {
		case SummaryField:
			summary = fmt.Sprintf("%s", e.Value)
		case DescriptionField:
			description = fmt.Sprintf("%s", e.Value)
		}
	}

	return &Issue{
		Key:    v.IssueKey,
		Status: v.CurrentStatus.Status,
		Title:  summary,
		Body:   description,
		Link:   v.Links.Web,
		Raw:    v,
	}
}
