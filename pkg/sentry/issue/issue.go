package issue

import "github.com/atlassian/go-sentry-api"

type Issue struct {
	Project string
	Type    string
	Title   string
	Link    string

	Raw *sentry.Issue
}
