package project

import "github.com/xanzy/go-gitlab"

func New(v *gitlab.Project) *Project {
	return &Project{
		Identifier: v.ID,
		Namespace:  v.Namespace.Path,
		Name:       v.Name,
		Link:       v.WebURL,
		Raw:        v,
	}
}
