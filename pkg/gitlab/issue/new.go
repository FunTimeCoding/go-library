package issue

import "github.com/xanzy/go-gitlab"

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
