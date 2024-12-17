package issue

import "gitlab.com/gitlab-org/api/client-go"

func New(v *gitlab.Issue) *Issue {
	return &Issue{
		Project:    v.ProjectID,
		Identifier: v.IID,
		Title:      v.Title,
		State:      v.State,
		Link:       v.WebURL,
		Raw:        v,
	}
}
