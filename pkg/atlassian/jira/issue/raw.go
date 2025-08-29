package issue

import "github.com/andygrunwald/go-jira"

func Raw(key string) *jira.Issue {
	result := RawStub()
	result.Key = key

	return result
}
