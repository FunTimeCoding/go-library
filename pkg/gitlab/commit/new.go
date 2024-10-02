package commit

import "github.com/xanzy/go-gitlab"

func New(v *gitlab.Commit) *Commit {
	return &Commit{
		Identifier: v.ID,
		Title:      v.Title,
		Message:    v.Message,
		Author:     v.AuthorName,
		Date:       v.CreatedAt,
		Link:       v.WebURL,
		Raw:        v,
	}
}
