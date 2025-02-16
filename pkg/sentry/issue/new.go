package issue

import "github.com/atlassian/go-sentry-api"

func New(v *sentry.Issue) *Issue {
	return &Issue{
		Project: v.Project.Name,
		Type:    *v.Type,
		Title:   *v.Title,
		Link:    *v.Permalink,
		Raw:     v,
	}
}
