package issue

import "github.com/atlassian/go-sentry-api"

func NewSlice(v []sentry.Issue) []*Issue {
	var result []*Issue

	for _, element := range v {
		result = append(result, New(&element))
	}

	return result
}
