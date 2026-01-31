package issue

import "github.com/atlassian/go-sentry-api"

func NewSlice(v []sentry.Issue) []*Issue {
	var result []*Issue

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
