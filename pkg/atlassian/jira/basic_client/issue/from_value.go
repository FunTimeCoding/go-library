package issue

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/customer"

func FromValue(v Values) *Issue {
	result := New()
	result.Key = v.IssueKey

	for _, e := range v.RequestFieldValues {
		switch e.FieldId {
		case customer.SummaryField:
			result.Title = e.Value.(string)
		case customer.DescriptionField:
			result.Body = e.Value.(string)
		}
	}

	return result
}
